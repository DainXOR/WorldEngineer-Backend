package test

import (
	"dainxor/we/base/logger"

	"github.com/gin-gonic/gin"
)

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
