package db

import (
	"dainxor/we/base/logger"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
	"net/http"
	"net/smtp"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
)

type mailType struct{}

var Mail mailType

func (mailType) SendErrorMail(email string) types.Optional[models.ErrorResponse] {
	mail := models.MailSend{}

	mail.Receiver(email)
	mail.Title("Error in verification")
	mail.MsgLine("Hey there!")
	mail.MsgWhiteLine()
	mail.MsgLine("We're sorry to inform you that there was an error in the verification process.")
	mail.MsgWhiteLine()
	mail.MsgLine("Don't worry, this doesn't mean you can't sign up later using this same email address.")
	mail.MsgLine("If you're still interested in signing up, please try again.")
	mail.MsgWhiteLine()
	mail.MsgLine("We apologize for the inconvenience.")
	mail.MsgLine("If you have any questions, feel free to contact us.")
	mail.MsgWhiteLine()
	mail.MsgLine("Best regards,")
	mail.MsgWhiteLine()
	mail.MsgLine("DainXOR")

	if err := Mail.SendEmail(mail); err.IsPresent() {
		return err
	}

	return types.OptionalEmpty[models.ErrorResponse]()
}

func (mailType) SendAuthMail(email string) types.Result[string, models.ErrorResponse] {
	resultCode := Auth.GenerateCode()
	if resultCode.IsErr() {
		return types.ResultErr[string](resultCode.Error())
	}

	code := resultCode.Value()

	mail := models.MailSend{}
	mail.Receiver(email)
	mail.Title("Verification")
	mail.MsgLine("Hey there! Thanks for signing up.")
	mail.MsgWhiteLine()
	mail.MsgLine("To continue, please verify your email.")
	mail.MsgLine("Use the following code to verify your email:")
	mail.MsgWhiteLine()
	mail.MsgSameLine("--------------------> ", code, " <--------------------")
	mail.MsgWhiteLine()
	mail.MsgLine("This code will expire in 5 minutes.")
	mail.MsgWhiteLine()
	mail.MsgLine("If you didn't sign up, ignore this email.")
	mail.MsgWhiteLine()
	mail.MsgLine("Have a nice day!")
	mail.MsgWhiteLine()
	mail.MsgWhiteLine()
	mail.MsgWhiteLine()
	mail.MsgLine("Best regards,")
	mail.MsgWhiteLine()
	mail.MsgLine("DainXOR")

	if err := Mail.SendEmail(mail); err.IsPresent() {
		err2, _ := utils.Retry(func() (types.Optional[models.ErrorResponse], error) { return Mail.SendErrorMail(email), nil },
			3,
			"Failed to send error email: ",
			"Could not send error email: ",
		)

		err := err.Get()
		if err2.IsPresent() {
			err2 := err2.Get()
			err.Code = types.Http.InternalServerError()
			err.Type += " | " + err2.Type
			err.Message += " | " + err2.Message
			err.Detail += " | " + err2.Detail

			return types.ResultErr[string](err)
		}

		return types.ResultErr[string](err)
	}

	return types.ResultOk[string, models.ErrorResponse](code)
}

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck().
		EnableDomainSuggest()
)

func (mailType) VerifyEmailAddress(email string) types.Optional[models.ErrorResponse] {
	logger.Info("Verifying email address: ", email)
	result, err := verifier.Verify(email)

	if err != nil {
		logger.Warning("Verify email address failed, error is: ", err.Error())

		return types.OptionalOf(
			models.Error(
				types.Http.BadRequest(),
				"failed",
				"could not verify the email address, please try again",
				err.Error(),
			),
		)
	}

	logger.Info("Email validation result", result)
	logger.Info(
		"Email:", result.Email,
		"\nReachable:", result.Reachable,
		"\nSyntax:", result.Syntax,
		"\nSMTP:", result.SMTP,
		"\nGravatar:", result.Gravatar,
		"\nSuggestion:", result.Suggestion,
		"\nDisposable:", result.Disposable,
		"\nRoleAccount:", result.RoleAccount,
		"\nFree:", result.Free,
		"\nHasMxRecords:", result.HasMxRecords)

	if !result.Syntax.Valid {
		logger.Warning("Invalid email address syntax")
		return types.OptionalOf(
			models.Error(
				http.StatusBadRequest,
				"invalid",
				"email address syntax is invalid",
			),
		)
	}
	if result.Disposable {
		logger.Warning("Disposable email address")
		return types.OptionalOf(
			models.Error(
				http.StatusBadRequest,
				"disposable",
				"disposable email addresses are not accepted",
			),
		)
	}

	// possible return string values: yes, no, unkown
	if result.Reachable == "no" {
		logger.Warning("Unreachable email address")

		if result.Suggestion != "" {
			logger.Warning("Suggestion:", result.Suggestion)
			return types.OptionalOf(
				models.Error(
					http.StatusBadRequest,
					"suggestion",
					"email address is not reachable",
					result.Suggestion,
				),
			)
		}

		return types.OptionalOf(
			models.Error(
				http.StatusBadRequest,
				"unreachable",
				"email address was unreachable",
			),
		)
	} else if result.Reachable == "unknown" {
		logger.Warning("Unknown email address reachability")
	}

	// check MX records so we know DNS setup properly to recieve emails
	if !result.HasMxRecords {
		logger.Warning("MX record not found")
		return types.OptionalOf(
			models.Error(
				http.StatusBadRequest,
				"mx",
				"domain entered not properly setup to recieve emails, MX record not found",
			),
		)
	}

	logger.Info("Email address is valid")
	return types.OptionalEmpty[models.ErrorResponse]()
}

var credentials models.SMTPCredentials

func (mailType) LoadCredentials() types.Optional[models.ErrorResponse] {
	username, ok1 := os.LookupEnv("SMTP_USERNAME")
	password, ok2 := os.LookupEnv("SMTP_PASSWORD")
	smtpHost, ok3 := os.LookupEnv("SMTP_HOST")
	smtpPort, ok4 := os.LookupEnv("SMTP_PORT")
	email, ok5 := os.LookupEnv("SMTP_EMAIL")

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
		return types.OptionalOf(models.Error(
			types.Http.InternalServerError(),
			"send_email",
			"Failed to send email",
			"SMTP credentials missing",
		))
	}

	credentials = models.SMTPCredentials{
		Username: username,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Email:    email,
	}

	return types.OptionalEmpty[models.ErrorResponse]()
}

func (mailType) SendEmail(mail models.MailSend) types.Optional[models.ErrorResponse] {
	username := credentials.Username
	password := credentials.Password
	smtpHost := credentials.SMTPHost
	smtpPort := credentials.SMTPPort
	email := credentials.Email

	auth := smtp.PlainAuth("", username, password, smtpHost)

	smtpUrl := smtpHost + ":" + smtpPort

	mail.Sender(email)
	logger.Info("Sending email from: ", mail.From)

	_, err := utils.Retry(
		func() (interface{}, error) {
			return nil, smtp.SendMail(smtpUrl, auth, mail.From, mail.To, mail.Message())
		},
		3,
		"Failed to send email: ",
		"Email could not be sent: ",
	)

	if err != nil {
		return types.OptionalOf(models.Error(
			types.Http.InternalServerError(),
			"send_email",
			"Failed to send email",
			err.Error(),
		))
	}

	return types.OptionalEmpty[models.ErrorResponse]()
}
