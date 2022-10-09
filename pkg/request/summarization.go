package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type Summarization struct {
	Request
	InputLanguage *string `validate:"omitempty,oneof=auto en es fr it pt ca"`
	// The following three fields are mutually exclusive
	Text     string `validate:"-"`
	URL      string `validate:"omitempty,url"`
	Document string `validate:"-"`
	// The following two fields are mutually exclusive
	Sentences string `validate:"omitempty,numeric"`
	Limit     string `validate:"omitempty,numeric"`
}

func (s *Summarization) Validate() error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if err := validateMutualExclusiveFields(
		map[string]string{
			"txt": s.Text,
			"url": s.URL,
			"doc": s.Document,
		},
	); err != nil {
		return err
	}

	return validateMutualExclusiveFields(
		map[string]string{
			"sentences": s.Sentences,
			"limit":     s.Limit,
		},
	)

}

func (s *Summarization) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", s.Key)
	multipartForm.AddOptionalField("lang", s.InputLanguage)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": s.Text,
			"url": s.URL,
			"doc": s.Document,
		},
	)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"sentences": s.Sentences,
			"limit":     s.Limit,
		},
	)

	return multipartForm, nil
}
