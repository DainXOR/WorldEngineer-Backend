package mail

import (
	"dainxor/we/base/logger"
	"dainxor/we/db"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
)

func SendErrorMail(email string) types.Optional[models.ErrorResponse] {
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

	if err := SendEmail(mail); err.IsPresent() {
		return err
	}

	return types.OptionalEmpty[models.ErrorResponse]()
}

func SendAuthMail(email string) types.Result[models.AuthCodeDB, models.ErrorResponse] {
	resultCode := db.Auth.GenerateCode()
	if resultCode.IsErr() {
		return types.ResultErr[models.AuthCodeDB](resultCode.Error())
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

	if err := SendEmail(mail); err.IsPresent() {
		err2, _ := utils.Retry(func() (types.Optional[models.ErrorResponse], error) { return SendErrorMail(email), nil },
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

			return types.ResultErr[models.AuthCodeDB](err)
		}

		return types.ResultErr[models.AuthCodeDB](err)
	}

	resultDB := db.Auth.SaveCode(email, code)
	if resultDB.IsErr() {
		err2, _ := utils.Retry(func() (types.Optional[models.ErrorResponse], error) { return SendErrorMail(email), nil },
			3,
			"Failed to send error email: ",
			"Could not send error email: ",
		)

		if db.Auth.GetValidCodeByEmail(email).IsOk() {
			logger.Info("Deleting code for email: ", email)
			db.Auth.DeleteAllCodesByEmail(email)
		} else {
			logger.Warning("Code not found for email: ", email)
		}

		err := resultDB.Error()
		if err2.IsPresent() {
			err2 := err2.Get()
			err.Code = types.Http.InternalServerError()
			err.Type += " | " + err2.Type
			err.Message += " | " + err2.Message
			err.Detail += " | " + err2.Detail

			return types.ResultErr[models.AuthCodeDB](err)
		}

		return types.ResultErr[models.AuthCodeDB](err)
	}

	return resultDB
}
