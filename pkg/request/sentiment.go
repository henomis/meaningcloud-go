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
	multipartForm.AddOptionalField("ilang", s.OutputLanguage)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": s.Text,
			"url": s.URL,
			"doc": s.Document,
		},
	)
	multipartForm.AddOptionalField("txtf", s.TextFormat)
	multipartForm.AddOptionalField("model", s.Model)
	multipartForm.AddOptionalField("verbose", s.Verbose)
	multipartForm.AddOptionalField("egp", s.ExpandGlobalPolarity)
	multipartForm.AddOptionalField("rt", s.ReliableText)
	multipartForm.AddOptionalField("uw", s.UnknownWords)
	multipartForm.AddOptionalField("dm", s.DisambiguationApplied)
	multipartForm.AddOptionalField("sdg", s.SemanticDisambiguationGrouping)
	multipartForm.AddOptionalField("cont", s.DisambiguationContext)
	multipartForm.AddOptionalField("ud", s.UserDictionary)

	return multipartForm, nil
}

func (s *Sentiment) String() string {
	bytes, _ := json.MarshalIndent(s, "", " ")
	return string(bytes)
}
