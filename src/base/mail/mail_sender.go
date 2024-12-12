package mail

import (
	"dainxor/we/base/logger"
	"dainxor/we/models"
	"dainxor/we/types"
	"dainxor/we/utils"
	"os"

	"net/smtp"

	"github.com/gin-gonic/gin"
)

var credentials models.SMTPCredentials

func LoadCredentials() types.Optional[models.ErrorResponse] {
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

func SendTestEmail1(*gin.Context) {
	logger.Warning("This test is disabled")

	/*
		mail := models.MailSend{}
		mail.Receiver("daniel.leond@udea.edu.co")
		mail.Receiver("dainxor@gmail.com")
		mail.Title("Hmmm")
		mail.MsgLine("Maybe use proton?")

		SendEmail(mail)
	*/
}

func SendTestEmail2(*gin.Context) {
	logger.Warning("This test is disabled")

	/*
		mail := models.MailSend{}
		mail.Receiver("dleon_2407@hotmail.com")
		mail.Receiver("dannyleon2001@hotmail.com")
		mail.Title("Hola mama :P")
		mail.MsgLine("Ando haciendo un trabajo y estoy intentando mandar correos de verificacion.")
		mail.MsgLine("Como lo esta leyendo, es que funciono :D")
		mail.MsgLine("Los quiero <3")

		SendEmail(mail)
	*/
}

func SendEmail(mail models.MailSend) types.Optional[models.ErrorResponse] {
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
