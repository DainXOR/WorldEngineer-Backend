package mail

import (
	"dainxor/we/auth"
	"dainxor/we/configs"
	"dainxor/we/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendAuthMail(c *gin.Context, email string) {
	if auth.HasCode(email) {
		c.JSON(http.StatusConflict,
			models.ErrorResponse{
				Type:    "conflict",
				Message: "Email is already in use",
				Detail:  "Check your email for the verification code, or try again in a few minutes",
				//Extra:   "Try looking in your spam folder",
			})

		return
	}

	mail := models.MailSend{}
	mail.Receiver(email)
	mail.Title("Verification")
	mail.MsgLine("Hey there! Thanks for signing up.")
	mail.MsgWhiteLine()
	mail.MsgLine("To continue, please verify your email.")
	mail.MsgLine("Use the following code to verify your email:")
	code := auth.GenerateCode()
	mail.MsgSameLine("--------------------> ", code, " <--------------------")
	mail.MsgWhiteLine()
	mail.MsgLine("If you didn't sign up, ignore this email.")
	mail.MsgWhiteLine()
	mail.MsgLine("Have a nice day!")
	mail.MsgWhiteLine()
	mail.MsgLine("Best regards,")
	mail.MsgLine("DainXOR")

	err := SendEmail(c, mail)

	if err == nil {
		// Save code to database for verification
		configs.DB.Create(&models.AuthCodeDB{
			Email:     email,
			Code:      code,
			CreatedAt: configs.DB.NowFunc(),
		})
	}
}
