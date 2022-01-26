package normalizer

import (
	"github.com/Masterminds/goutils"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strings"
)

func Email(email string) string {
	return strings.ToLower(email)
}

func PhoneNumber(phone string) string {
	parsed, err := phonenumbers.Parse(phone, "DE")
	if err != nil {
		return phone
	}
	return phonenumbers.Format(parsed, phonenumbers.INTERNATIONAL)
}

func Name(name string) string {
	return normalizeWhitespace(goutils.CapitalizeFully(name, ' ', '-'))
}

func normalizeWhitespace(text string) string {
	space, err := regexp.Compile(`\s+`)
	if err != nil {
		return text
	}
	duplicateRemoved := space.ReplaceAllString(text, " ")
	return strings.Trim(duplicateRemoved, " ")
}
