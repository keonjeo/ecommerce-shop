package locale

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// T returns the translated msg identified by messageID
// second argument is optional map[string]interface{}{} or interface{} used as a template
var T TranslateFunc

// TP returns the translated msg identified by messageID
var TP TranslatePluralFunc

// TranslateFunc returns the translated msg identified by messageID
// second argument is optional map[string]interface{}{} or struct{} used as a template
type TranslateFunc func(messageID string, args ...interface{}) string

// TranslatePluralFunc returns the translated msg identified by messageID
// pluralCount can be (int, int8, int16, int32, int64) or a float formatted as a string (e.g. "123.45")
// if no template is passed (TemplateData is nil), pluralCount will be used as a template
// third argument is optional map[string]interface{}{} or struct{} used as a template
type TranslatePluralFunc func(messageID string, pluralCount interface{}, args ...interface{}) string

// locales holds the list of all the locales
var locales map[string]string = make(map[string]string)

// localizers holds the list of all i18n localizers that can be fetched at each request
var localizers map[string]*i18n.Localizer = make(map[string]*i18n.Localizer)

// CreateBundle creates the default i18n bundle
func CreateBundle() *i18n.Bundle {
	return i18n.NewBundle(language.English)
}

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

// InitLocales initializes the localizers and stores them in map
func InitLocales(b *i18n.Bundle) {
	en := i18n.NewLocalizer(b, "en")
	sr := i18n.NewLocalizer(b, "sr")

	localizers["en"] = en
	localizers["sr"] = sr
}

// GetLocalizer gets the localizer by the specified language key
func GetLocalizer(lang string) (*i18n.Localizer, error) {
	val, ok := localizers[lang]
	if !ok {
		return nil, errors.New("language with the given key is not supported")
	}
	return val, nil
}

// GetSupportedLocales shows the list of supported locales
func GetSupportedLocales() map[string]string {
	return locales
}

// InitTranslations configures i18n
func InitTranslations() {
	bundle := CreateBundle()
	LoadTranslations(bundle)
	InitLocales(bundle)

	T = TFuncWithLanguage("en")
	TP = TPFuncWithLanguage("en")
}

// TFuncWithLanguage returns the TranslateFunc with specific language preference
func TFuncWithLanguage(pref string) TranslateFunc {
	localizer, _ := GetLocalizer(pref)
	return func(messageID string, args ...interface{}) string {
		lc := &i18n.LocalizeConfig{
			MessageID: messageID,
		}
		if len(args) == 1 {
			lc.TemplateData = args[0]
		}
		return localizer.MustLocalize(lc)
	}
}

// TPFuncWithLanguage returns the TranslatePluralFunc with specific language preference
func TPFuncWithLanguage(pref string) TranslatePluralFunc {
	localizer, _ := GetLocalizer(pref)
	return func(messageID string, pluralCount interface{}, args ...interface{}) string {
		lc := &i18n.LocalizeConfig{
			MessageID:   messageID,
			PluralCount: pluralCount,
		}
		if len(args) == 1 {
			data := args[0].(map[string]interface{})
			data["PluralCount"] = fmt.Sprintf("%v", pluralCount)
			lc.TemplateData = data
		}
		return localizer.MustLocalize(lc)
	}
}
