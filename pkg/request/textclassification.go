package request

import (
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

	return validateMutualExclusiveFields(
		map[string]string{
			"txt": t.Text,
			"url": t.URL,
			"doc": t.Document,
		},
	)

}

func (t *TextClassification) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := t.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", t.Key)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": t.Text,
			"url": t.URL,
			"doc": t.Document,
		},
	)
	multipartForm.AddField("model", t.Model)
	multipartForm.AddOptionalField("verbose", t.Verbose)
	multipartForm.AddOptionalField("title", t.Title)
	multipartForm.AddOptionalField("abstract", t.Abstract)
	multipartForm.AddOptionalField("categories_filter", t.CategoriesFilter)
	multipartForm.AddOptionalField("expand_hierarchy", t.ExpandHierarchy)

	return multipartForm, nil
}
