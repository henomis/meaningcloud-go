package request

import (
	"encoding/json"
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

	if s.Text == "" && s.URL == "" && s.Document == "" {
		return fmt.Errorf("one of the following fields must be set: txt, url, doc")
	}

	if s.Text != "" && (s.URL != "" || s.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if s.URL != "" && (s.Text != "" || s.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if s.Document != "" && (s.Text != "" || s.URL != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if s.Sentences != "" && s.Limit != "" {
		return fmt.Errorf("only one of the following fields can be set: sentences, limit")
	}

	if s.Sentences == "" && s.Limit == "" {
		return fmt.Errorf("one of the following fields must be set: sentences, limit")
	}

	return nil

}

func (s *Summarization) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", s.Key)

	if s.InputLanguage != nil {
		multipartForm.AddField("lang", *s.InputLanguage)
	}

	if s.Text != "" {
		multipartForm.AddField("txt", s.Text)
	} else if s.URL != "" {
		multipartForm.AddField("url", s.URL)
	} else if s.Document != "" {
		multipartForm.AddFile("doc", s.Document)
	}

	if s.Sentences != "" {
		multipartForm.AddField("sentences", s.Sentences)
	} else if s.Limit != "" {
		multipartForm.AddField("limit", s.Limit)
	}

	return multipartForm, nil
}

func (s *Summarization) String() string {
	bytes, _ := json.MarshalIndent(s, "", " ")
	return string(bytes)
}
