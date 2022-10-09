package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type LanguageIdentification struct {
	Request
	// The following tow fields are mutually exclusive
	Text string `validate:"-"`
	URL  string `validate:"omitempty,url"`
}

func (l *LanguageIdentification) Validate() error {
	validate := validator.New()

	if err := validate.Struct(l); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if l.Text == "" && l.URL == "" {
		return fmt.Errorf("one of the following fields must be set: txt, url, doc")
	}

	if l.Text != "" && l.URL != "" {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	return nil

}

func (l *LanguageIdentification) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := l.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", l.Key)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": l.Text,
			"url": l.URL,
		},
	)

	return multipartForm, nil
}

func (l *LanguageIdentification) String() string {
	bytes, _ := json.MarshalIndent(l, "", " ")
	return string(bytes)
}
