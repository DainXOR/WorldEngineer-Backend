package mail

import (
	"dainxor/we/logger"
	"dainxor/we/models"
	"dainxor/we/utils"

	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func SendTestEmail1(c *gin.Context) {

	logger.Warning("This test is disabled")
	return

	mail := models.MailSend{}
	mail.Receiver("daniel.leond@udea.edu.co")
	mail.Receiver("dainxor@gmail.com")
	mail.Title("Hmmm")
	mail.MsgLine("Maybe use proton?")

	SendEmail(c, mail)
}

func SendTestEmail2(c *gin.Context) {

	logger.Warning("This test is disabled")
	return

	mail := models.MailSend{}
	mail.Receiver("dleon_2407@hotmail.com")
	mail.Receiver("dannyleon2001@hotmail.com")
	mail.Title("Hola mama :P")
	mail.MsgLine("Ando haciendo un trabajo y estoy intentando mandar correos de verificacion.")
	mail.MsgLine("Como lo esta leyendo, es que funciono :D")
	mail.MsgLine("Los quiero <3")

	SendEmail(c, mail)
}

func SendEmail(c *gin.Context, mail models.MailSend) error {
	username := "api"
	password := "20a57210c28ca0c202f5192e7f06d339"
	smtpHost := "live.smtp.mailtrap.io"
	smtpPort := "587"

	auth := smtp.PlainAuth("", username, password, smtpHost)

	smtpUrl := smtpHost + ":" + smtpPort

	mail.Sender("sign.we@fardina143.co")
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
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "unable to send mail to the address, please try again",
			},
		)
	}

	return err
}
