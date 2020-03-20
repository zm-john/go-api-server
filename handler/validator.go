package handler

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni *ut.UniversalTranslator

	// Validate global
	Validate *validator.Validate

	// Trans global
	Trans ut.Translator
)

// ValidateErrorToErrResponse 将验证错误转换为响应错误
func ValidateErrorToErrResponse(err error) []Error {
	errs := err.(validator.ValidationErrors)

	responseErrs := make([]Error, 0)
	for _, val := range errs {
		responseErrs = append(responseErrs, Error{
			Field:   val.Field(),
			Message: val.Translate(Trans),
		})
	}

	return responseErrs
}

func init() {
	zh := zh.New()
	en := en.New()
	uni = ut.New(en, zh)
	Trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	zh_translations.RegisterDefaultTranslations(Validate, Trans)
}
