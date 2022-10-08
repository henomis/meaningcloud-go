package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type DeepCategorization struct {
	Request
	// The following three fields are mutually exclusive
	Text           string  `validate:"-"`
	URL            string  `validate:"omitempty,url"`
	Document       string  `validate:"-"`
	Model          string  `validate:"-"`
	Verbose        *string `validate:"omitempty,oneof=y n"`
	Popolarity     *string `validate:"omitempty,oneof=y n"`
	UserDictionary *string `validate:"-"`
}

func (d *DeepCategorization) Validate() error {
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

func (d *DeepCategorization) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := d.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", d.Key)

	if d.Text != "" {
		multipartForm.AddField("txt", d.Text)
	} else if d.URL != "" {
		multipartForm.AddField("url", d.URL)
	} else if d.Document != "" {
		multipartForm.AddFile("doc", d.Document)
	}

	multipartForm.AddField("model", d.Model)

	if d.Verbose != nil {
		multipartForm.AddField("verbose", *d.Verbose)
	}
	if d.Popolarity != nil {
		multipartForm.AddField("popolarity", *d.Popolarity)
	}
	if d.UserDictionary != nil {
		multipartForm.AddField("ud", *d.UserDictionary)
	}

	return multipartForm, nil
}

func (d *DeepCategorization) String() string {
	bytes, _ := json.MarshalIndent(d, "", " ")
	return string(bytes)
}
