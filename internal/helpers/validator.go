package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translator "github.com/go-playground/validator/v10/translations/ru"
	"github.com/sirupsen/logrus"
)

var validate22 *validator.Validate = validator.New()

func ValidateEmail(s string) error {
	err := validate22.Var(s, "required,email")
	if err != nil {
		return fmt.Errorf("[email:%s] email is not valid", s)
	}
	return nil
}

func ValidateOptionalEmail(s string) error {
	err := validate22.Var(s, "email")
	if err != nil {
		return fmt.Errorf("[email:%s] email is not valid", s)
	}
	return nil
}

func ValidateColor(color string) error {
	if len(color) != 7 {
		return fmt.Errorf("[color:%s] цвет должен быть #000000 (#xxxxxx)", color)
	}
	if color[0] != '#' {
		return fmt.Errorf("[color:%s] цвет должен начинаться с # (#xxxxxx)", color)
	}
	if _, err := strconv.ParseInt(color[1:], 16, 64); err != nil {
		return fmt.Errorf("[color:%s] цвет может содержать только цифры #000000 (#xxxxxx)", color)
	}
	return nil
}

// HTTPSValidation — строка должна начинаться с https://
func HTTPSValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return strings.HasPrefix(value, "https://")
}

// TrimValidation — нет ведущих/замыкающих пробелов
func TrimValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return !(strings.HasPrefix(value, " ") || strings.HasSuffix(value, " "))
}

// ColorValidation — формат #RRGGBB
func ColorValidation(fl validator.FieldLevel) bool {
	color := fl.Field().String()
	if len(color) != 7 {
		return false
	}
	if color[0] != '#' {
		return false
	}
	if _, err := strconv.ParseInt(color[1:], 16, 64); err != nil {
		return false
	}
	return true
}

// NameValidation — только буквы/цифры/пробелы/( ) - _
func NameValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	rgxp := regexp.MustCompile(`^[\p{L}0-9\s()\-_]*$`)
	if !rgxp.MatchString(value) {
		panic(errors.New("название должно содержать только буквы, цифры и пробелы"))
	}
	return !(strings.HasPrefix(value, " ") || strings.HasSuffix(value, " "))
}

// ValidateLegalEntityField — число, не начинающееся с 2+ нулей
func ValidateLegalEntityField(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if !regexp.MustCompile(`^\d+$`).MatchString(value) {
		return false
	}
	return !regexp.MustCompile(`^0{2,}`).MatchString(value)
}

// OptionalEmailValidation — пусто или корректный email
func OptionalEmailValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if strings.TrimSpace(value) == "" {
		return true
	}
	return validate22.Var(value, "email") == nil
}

// ValidationStruct — валидация структуры с RU-переводами и кастомными правилами
func ValidationStruct[T any](str T, fields ...string) ([]string, bool) {
	russian := ru.New()
	uni := ut.New(russian, russian)

	trans, found := uni.GetTranslator("ru")
	if !found {
		logrus.Error("translator not found")
	}

	validate := validator.New()

	// Кастомные правила
	if err := validate.RegisterValidation("is_https", HTTPSValidation); err != nil {
		logrus.Error(err)
	}
	if err := validate.RegisterValidation("trim", TrimValidation); err != nil {
		logrus.Error(err)
	}
	if err := validate.RegisterValidation("color", ColorValidation); err != nil {
		logrus.Error(err)
	}
	if err := validate.RegisterValidation("name", NameValidation); err != nil {
		logrus.Error(err)
	}
	if err := validate.RegisterValidation("legal_entity_field", ValidateLegalEntityField); err != nil {
		logrus.Error(err)
	}
	if err := validate.RegisterValidation("optional_email", OptionalEmailValidation); err != nil {
		logrus.Error(err)
	}

	// Переводы кастомных правил
	if err := validate.RegisterTranslation("trim", trans,
		func(ut ut.Translator) error { return ut.Add("trim", "нужно убрать пробелы", true) },
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("trim", fe.Field(), fe.Param())
			if err != nil {
				logrus.Error(err)
			}
			return t
		},
	); err != nil {
		logrus.Error(err)
	}

	if err := validate.RegisterTranslation("color", trans,
		func(ut ut.Translator) error {
			return ut.Add("color", "цвет должен быть #000000 - #999999", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("color", fe.Field(), fe.Param())
			if err != nil {
				logrus.Error(err)
			}
			return t
		},
	); err != nil {
		logrus.Error(err)
	}

	if err := validate.RegisterTranslation("name", trans,
		func(ut ut.Translator) error { return ut.Add("name", "название а-Я и цифры 0-9", true) },
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("name", fe.Field(), fe.Param())
			if err != nil {
				logrus.Error(err)
			}
			return t
		},
	); err != nil {
		logrus.Error(err)
	}

	if err := validate.RegisterTranslation("legal_entity_field", trans,
		func(ut ut.Translator) error {
			return ut.Add("legal_entity", "{0} должно быть числом и начинаться с не более чем одного нуля", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("legal_entity", fe.Field(), fe.Param())
			if err != nil {
				logrus.Error(err)
			}
			return t
		},
	); err != nil {
		logrus.Error(err)
	}

	// Базовые RU-переводы
	if err := ru_translator.RegisterDefaultTranslations(validate, trans); err != nil {
		logrus.Error(err)
	}

	// Доп. перевод email
	if err := validate.RegisterTranslation("email", trans,
		func(ut ut.Translator) error {
			return ut.Add("email", "Почта должна быть действительной", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T("email", fe.Field(), fe.Param())
			if err != nil {
				logrus.Error(err)
			}
			return t
		},
	); err != nil {
		logrus.Error(err)
	}

	// Отображаем RU-названия полей через тег `ru:"..."`.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("ru"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Валидация
	var verr error
	if len(fields) > 0 {
		verr = validate.StructPartial(str, fields...)
	} else {
		verr = validate.Struct(str)
	}

	if verr != nil {
		var errs validator.ValidationErrors
		if errors.As(verr, &errs) {
			out := make([]string, 0, len(errs))
			for _, e := range errs {
				out = append(out, e.Translate(trans))
			}
			return out, false
		}
		return []string{verr.Error()}, false
	}

	return []string{}, true
}
