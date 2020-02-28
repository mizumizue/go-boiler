package validator

import (
	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja "github.com/go-playground/validator/v10/translations/ja"
)

type Validator struct {
	v  *validator.Validate
	ut ut.Translator
}

func New() *Validator {
	v := validator.New()
	trans := ut.New(ja_JP.New())
	jaTrans, _ := trans.GetTranslator("ja_JP")
	err := ja.RegisterDefaultTranslations(v, jaTrans)
	if err != nil {
		panic(err)
	}
	return &Validator{
		v,
		jaTrans,
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.v.Struct(i)
}

func (v *Validator) ValidationStrings(err error) []string {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("passed not validator error. arg should be validator.ValidationErrors.")
	}
	errs := make([]string, 0, len(validationErrs))
	for _, validationErr := range validationErrs {
		errs = append(errs, validationErr.Translate(v.ut))
	}
	return errs
}
