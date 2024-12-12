package mail

import (
	"dainxor/we/base/logger"
	"dainxor/we/models"
	"dainxor/we/types"

	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck().
		EnableDomainSuggest()
)

func VerifyEmailAddress(email string) types.Optional[models.ErrorResponse] {
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
