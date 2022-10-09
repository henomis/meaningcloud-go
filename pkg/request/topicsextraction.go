package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type TopicsExtraction struct {
	Request
	InputLanguage  string  `validate:"required,oneof=auto en es fr it pt ca"`
	OutputLanguage *string `validate:"omitempty,oneof=auto en es fr it pt ca"`
	// The following three fields are mutually exclusive
	Text                           string  `validate:"-"`
	URL                            string  `validate:"omitempty,url"`
	Document                       string  `validate:"-"`
	TextFormat                     *string `validate:"omitempty,oneof=plain markup"`
	TopicTypes                     *string `validate:"-"`
	UnknownWords                   *string `validate:"omitempty,oneof=y n"`
	RelaxedTypography              *string `validate:"omitempty,oneof=y u n"`
	UserDictionary                 *string `validate:"-"`
	ShowSubtopics                  *string `validate:"omitempty,oneof=y n"`
	DisambiguationApplied          *string `validate:"omitempty,oneof=n m s"`
	SemanticDisambiguationGrouping *string `validate:"omitempty,oneof=n g t l"`
	DisambiguationContext          *string `validate:"-"`
	TimeReference                  *string `validate:"-"`
}

func (t *TopicsExtraction) Validate() error {
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

func (t *TopicsExtraction) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := t.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", t.Key)
	multipartForm.AddField("lang", t.InputLanguage)
	if t.OutputLanguage != nil {
		multipartForm.AddField("ilang", *t.OutputLanguage)
	}

	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": t.Text,
			"url": t.URL,
			"doc": t.Document,
		},
	)

	multipartForm.AddOptionalField("txtf", t.TextFormat)
	multipartForm.AddOptionalField("tt", t.TopicTypes)
	multipartForm.AddOptionalField("uw", t.UnknownWords)
	multipartForm.AddOptionalField("rt", t.RelaxedTypography)
	multipartForm.AddOptionalField("ud", t.UserDictionary)
	multipartForm.AddOptionalField("st", t.ShowSubtopics)
	multipartForm.AddOptionalField("dm", t.DisambiguationApplied)
	multipartForm.AddOptionalField("sdg", t.SemanticDisambiguationGrouping)
	multipartForm.AddOptionalField("cont", t.DisambiguationContext)
	multipartForm.AddOptionalField("timeref", t.TimeReference)

	return multipartForm, nil
}
