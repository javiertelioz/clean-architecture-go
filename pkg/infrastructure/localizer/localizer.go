package localizer

import (
	_ "github.com/javiertelioz/clean-architecture-go/pkg/presentation/i18n"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var defaultLocale = "en-us"

type Localizer struct {
	ID      string
	printer *message.Printer
}

var locales = map[string]Localizer{
	"en-us": {
		ID:      "en-us",
		printer: message.NewPrinter(language.MustParse("en-US")),
	},
	"es-mx": {
		ID:      "es-mx",
		printer: message.NewPrinter(language.MustParse("es-MX")),
	},
}

func Get(language string) (Localizer, bool) {
	locale, exists := locales[strings.ToLower(language)]

	if !exists {
		return locales[defaultLocale], false
	}

	return locale, true
}

func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
