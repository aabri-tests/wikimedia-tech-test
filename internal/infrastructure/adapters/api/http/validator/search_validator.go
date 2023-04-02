package validator

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/wikimedia/internal/utils"
)

var (
	AllowedLanguages = []string{"en", "ar", "de", "fr", "it"}
)

type SearchQuery struct {
	Keyword  string `form:"keyword" binding:"required"`
	Language string `form:"language"`
}

func ValidateSearchRequest(q SearchQuery) error {
	if err := validateRequiredFields(q); err != nil {
		return err
	}
	if !utils.InArray(q.Language, AllowedLanguages) {
		return errors.New(fmt.Sprintf("Only %#v are supported", AllowedLanguages))
	}
	return nil
}

func validateRequiredFields(q SearchQuery) error {
	if q.Keyword == "" {
		return errors.New("missing required parameter 'keyword'")
	}
	return nil
}

func IsLanguageAllowed(language string) bool {
	for _, allowedLanguage := range AllowedLanguages {
		if allowedLanguage == language {
			return true
		}
	}
	return false
}
