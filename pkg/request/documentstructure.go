package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type DocumentStructure struct {
	Request
	// The following three fields are mutually exclusive
	Text     string `validate:"-"`
	URL      string `validate:"omitempty,url"`
	Document string `validate:"-"`
}

func (d *DocumentStructure) Validate() error {
	validate := validator.New()

	if err := validate.Struct(d); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if d.Text == "" && d.URL == "" && d.Document == "" {
		return fmt.Errorf("one of the following fields must be set: txt, url, doc")
	}

	if d.Text != "" && (d.URL != "" || d.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if d.URL != "" && (d.Text != "" || d.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if d.Document != "" && (d.Text != "" || d.URL != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	return nil

}

func (d *DocumentStructure) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := d.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", d.Key)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": d.Text,
			"url": d.URL,
			"doc": d.Document,
		},
	)

	return multipartForm, nil
}
