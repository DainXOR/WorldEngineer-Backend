package mail

import (
	"dainxor/we/logger"
	"dainxor/we/models"

	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck().
		EnableDomainSuggest()
)

func VerifyEmailAddress(c *gin.Context, email string) bool {
	logger.Info("Verifying email address: ", email)
	result, err := verifier.Verify(email)
	if err != nil {
		logger.Warning("Verify email address failed, error is: ", err)
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{
				Type:    "failed",
				Message: "could not verify the email address, please try again",
			},
		)
		return false
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
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:   "invalid",
				Detail: "email address syntax is invalid",
			},
		)
		return false
	}
	if result.Disposable {
		logger.Warning("Disposable email address")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:    "disposable",
				Message: "disposable email addresses are not accepted",
			},
		)
		return false
	}
	if result.Suggestion != "" {
		logger.Warning("Unreachable email address")
		logger.Warning("Suggestion: ", result.Suggestion)
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:    "suggestion",
				Message: "email address is not reachable",
				Detail:  result.Suggestion,
			},
		)
		return false
	}
	// possible return string values: yes, no, unkown
	if result.Reachable == "no" {
		logger.Warning("Unreachable email address")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:   "unreachable",
				Detail: "email address was unreachable",
			},
		)
		return false
	} else if result.Reachable == "unknown" {
		logger.Warning("Unknown email address reachability")
	}

	// check MX records so we know DNS setup properly to recieve emails
	if !result.HasMxRecords {
		logger.Warning("MX record not found")
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{
				Type:   "mx",
				Detail: "domain entered not properly setup to recieve emails, MX record not found",
			},
		)
		return false
	}

	logger.Info("Email address is valid")
	return true
}
