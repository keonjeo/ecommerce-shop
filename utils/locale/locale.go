package locale

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const DEFAULT_LOCALE = "en"

// T translates the msg identified by the id
var T TranslateFunc

// TranslateFunc returns the translated msg identified by messageID
type TranslateFunc func(messageID string, pluralCount interface{}, template interface{}) string

// locales holds the list of all the locales
var locales map[string]string = make(map[string]string)

// localizers holds the list of all i18n localizers that can be fetched at each request
var localizers map[string]*i18n.Localizer = make(map[string]*i18n.Localizer)

// LoadTranslations loads the translation files from the locales directory
func LoadTranslations(b *i18n.Bundle) {
	b.RegisterUnmarshalFunc("json", json.Unmarshal)
	localesDir := fmt.Sprintf("./locales")
	files, _ := ioutil.ReadDir(localesDir)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = filepath.Join(localesDir, filename)
			b.MustLoadMessageFile(fmt.Sprintf("./") + filepath.Join(localesDir, filename))
		}
	}
}

// GetLocalizerWithFallback gets the localizer by the specified language key
func GetLocalizerWithFallback(lang string) *i18n.Localizer {
	val, ok := localizers[lang]
	if !ok {
		val = localizers[DEFAULT_LOCALE]
	}
	return val
}

// GetSupportedLocales shows the list of supported locales
func GetSupportedLocales() map[string]string {
	return locales
}

// InitTranslations configures i18n
func InitTranslations() {
	bundle := i18n.NewBundle(language.English)
	LoadTranslations(bundle)
	localizers["en"] = i18n.NewLocalizer(bundle, "en")
	localizers["sr"] = i18n.NewLocalizer(bundle, "sr")
	T = TFuncWithLanguage("en")
}

// TFuncWithLanguage returns the TranslateFunc with specific language preference
func TFuncWithLanguage(pref string) TranslateFunc {
	localizer := GetLocalizerWithFallback(pref)
	return func(messageID string, pluralCount interface{}, template interface{}) string {
		lc := &i18n.LocalizeConfig{
			MessageID:    messageID,
			PluralCount:  pluralCount,
			TemplateData: template,
		}
		return localizer.MustLocalize(lc)
	}
}

// GetUserTranslations gets T func by the given locale
func GetUserTranslations(locale string) TranslateFunc {
	if _, ok := locales[locale]; !ok {
		locale = DEFAULT_LOCALE
	}
	return TFuncWithLanguage(locale)
}

// GetTranslationsAndLocale gets T fun together with locale from req headers
func GetTranslationsAndLocale(w http.ResponseWriter, r *http.Request) (TranslateFunc, string) {
	headerLocaleFull := strings.Split(r.Header.Get("Accept-Language"), ",")[0]
	headerLocale := strings.Split(strings.Split(r.Header.Get("Accept-Language"), ",")[0], "-")[0]
	defaultLocale := DEFAULT_LOCALE
	if locales[headerLocaleFull] != "" {
		return TFuncWithLanguage(headerLocaleFull), headerLocaleFull
	} else if locales[headerLocale] != "" {
		return TFuncWithLanguage(headerLocale), headerLocale
	} else if locales[defaultLocale] != "" {
		return TFuncWithLanguage(defaultLocale), headerLocale
	}
	return TFuncWithLanguage(defaultLocale), defaultLocale
}
