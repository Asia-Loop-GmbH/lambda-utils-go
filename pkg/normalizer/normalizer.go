package normalizer

import (
	"regexp"
	"strings"

	"github.com/Masterminds/goutils"
	"github.com/nyaruka/phonenumbers"
	"github.com/sirupsen/logrus"
)

func Email(log *logrus.Entry, email string) string {
	result := strings.ToLower(email)
	log.Infof("normalized email: '%s' -> '%s'", email, result)
	return result
}

func PhoneNumber(log *logrus.Entry, phone string) string {
	parsed, err := phonenumbers.Parse(phone, "DE")
	if err != nil {
		log.Warnf("could not normalize phone '%s', return original", phone)
		return phone
	}
	result := phonenumbers.Format(parsed, phonenumbers.INTERNATIONAL)
	log.Infof("normalized phone: '%s' -> '%s'", phone, result)
	return result
}

func Name(log *logrus.Entry, name string) string {
	result := normalizeWhitespace(goutils.CapitalizeFully(name, ' ', '-'))
	log.Infof("normalized name: '%s' -> '%s'", name, result)
	return result
}

func normalizeWhitespace(text string) string {
	space, err := regexp.Compile(`\s+`)
	if err != nil {
		return text
	}
	duplicateRemoved := space.ReplaceAllString(text, " ")
	return strings.Trim(duplicateRemoved, " ")
}
