package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type TextClassification struct {
	Request
	// The following three fields are mutually exclusive
	Text             string  `validate:"-"`
	URL              string  `validate:"omitempty,url"`
	Document         string  `validate:"-"`
	Model            string  `validate:"-"`
	Verbose          *string `validate:"omitempty,oneof=y n"`
	Title            *string `validate:"-"`
	Abstract         *string `validate:"-"`
	CategoriesFilter *string `validate:"-"`
	ExpandHierarchy  *string `validate:"omitempty,oneof=n p a"`
}

func (t *TextClassification) Validate() error {
	validate := validator.New()

	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if t.Text == "" && t.URL == "" && t.Document == "" {
		return fmt.Errorf("one of the following fields must be set: txt, url, doc")
	}

	if t.Text != "" && (t.URL != "" || t.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if t.URL != "" && (t.Text != "" || t.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if t.Document != "" && (t.Text != "" || t.URL != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	return nil

}

func (t *TextClassification) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := t.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", t.Key)

	if t.Text != "" {
		multipartForm.AddField("txt", t.Text)
	} else if t.URL != "" {
		multipartForm.AddField("url", t.URL)
	} else if t.Document != "" {
		multipartForm.AddFile("doc", t.Document)
	}

	multipartForm.AddField("model", t.Model)

	if t.Verbose != nil {
		multipartForm.AddField("verbose", *t.Verbose)
	}
	if t.Title != nil {
		multipartForm.AddField("title", *t.Title)
	}
	if t.Abstract != nil {
		multipartForm.AddField("abstract", *t.Abstract)
	}
	if t.CategoriesFilter != nil {
		multipartForm.AddField("categories_filter", *t.CategoriesFilter)
	}
	if t.ExpandHierarchy != nil {
		multipartForm.AddField("expand_hierarchy", *t.ExpandHierarchy)
	}

	return multipartForm, nil
}

func (t *TextClassification) String() string {
	bytes, _ := json.MarshalIndent(t, "", " ")
	return string(bytes)
}
