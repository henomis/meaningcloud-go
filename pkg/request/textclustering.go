package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type TextClustering struct {
	Request
	InputLanguage string  `validate:"required,oneof=auto en es fr it pt ca da sv no fi zh ru ar"`
	Text          string  `validate:"-"`
	ID            *string `validate:"-"`
	Mode          *string `validate:"omitempty,oneof=tm dg"`
	StopWords     *string `validate:"-"`
}

func (t *TextClustering) Validate() error {

	validate := validator.New()

	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if t.Text == "" {
		return fmt.Errorf("the following fields must be set: txt")
	}

	return nil

}

func (t *TextClustering) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := t.Validate(); err != nil {
		return nil, err
	}

	multipartform := multipartform.New()

	multipartform.AddField("key", t.Key)
	multipartform.AddField("lang", t.InputLanguage)
	multipartform.AddField("txt", t.Text)
	multipartform.AddOptionalField("id", t.ID)
	multipartform.AddOptionalField("mode", t.Mode)
	multipartform.AddOptionalField("sw", t.StopWords)

	return multipartform, nil

}
