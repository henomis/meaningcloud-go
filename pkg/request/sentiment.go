package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type Sentiment struct {
	Request
	InputLanguage  string  `validate:"required,oneof=auto en es fr it pt ca"`
	OutputLanguage *string `validate:"omitempty,oneof=auto en es fr it pt ca"`
	// The following three fields are mutually exclusive
	Text                           string  `validate:"-"`
	URL                            string  `validate:"omitempty,url"`
	Document                       string  `validate:"-"`
	TextFormat                     *string `validate:"omitempty,oneof=plain markup"`
	Model                          *string `validate:"omitempty,oneof=general"`
	Verbose                        *string `validate:"omitempty,oneof=y n"`
	ExpandGlobalPolarity           *string `validate:"omitempty,oneof=y n"`
	ReliableText                   *string `validate:"omitempty,oneof=y u n"`
	UnknownWords                   *string `validate:"omitempty,oneof=y n"`
	DisambiguationApplied          *string `validate:"omitempty,oneof=n m s"`
	SemanticDisambiguationGrouping *string `validate:"omitempty,oneof=n g t l"`
	DisambiguationContext          *string `validate:"-"`
	UserDictionary                 *string `validate:"-"`
}

func (s *Sentiment) Validate() error {
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

	return nil

}

func (s *Sentiment) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", s.Key)
	multipartForm.AddField("lang", s.InputLanguage)
	if s.OutputLanguage != nil {
		multipartForm.AddField("ilang", *s.OutputLanguage)
	}

	if s.Text != "" {
		multipartForm.AddField("txt", s.Text)
	} else if s.URL != "" {
		multipartForm.AddField("url", s.URL)
	} else if s.Document != "" {
		multipartForm.AddFile("doc", s.Document)
	}

	if s.TextFormat != nil {
		multipartForm.AddField("txtf", *s.TextFormat)
	}
	if s.Model != nil {
		multipartForm.AddField("model", *s.Model)
	}
	if s.Verbose != nil {
		multipartForm.AddField("verbose", *s.Verbose)
	}
	if s.ExpandGlobalPolarity != nil {
		multipartForm.AddField("egp", *s.ExpandGlobalPolarity)
	}
	if s.ReliableText != nil {
		multipartForm.AddField("rt", *s.ReliableText)
	}
	if s.UnknownWords != nil {
		multipartForm.AddField("uw", *s.UnknownWords)
	}
	if s.DisambiguationApplied != nil {
		multipartForm.AddField("dm", *s.DisambiguationApplied)
	}
	if s.SemanticDisambiguationGrouping != nil {
		multipartForm.AddField("sdg", *s.SemanticDisambiguationGrouping)
	}
	if s.DisambiguationContext != nil {
		multipartForm.AddField("cont", *s.DisambiguationContext)
	}
	if s.UserDictionary != nil {
		multipartForm.AddField("ud", *s.UserDictionary)
	}

	return multipartForm, nil
}

func (s *Sentiment) String() string {
	bytes, _ := json.MarshalIndent(s, "", " ")
	return string(bytes)
}
