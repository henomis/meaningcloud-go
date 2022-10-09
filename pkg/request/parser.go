package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type Parser struct {
	Request
	InputLanguage  string  `validate:"required,oneof=auto en es fr it pt ca"`
	OutputLanguage *string `validate:"omitempty,oneof=auto en es fr it pt ca"`
	// The following three fields are mutually exclusive
	Text                           string  `validate:"-"`
	URL                            string  `validate:"omitempty,url"`
	Document                       string  `validate:"-"`
	Verbose                        *string `validate:"omitempty,oneof=y n"`
	TextFormat                     *string `validate:"omitempty,oneof=plain markup"`
	UnknownWords                   *string `validate:"omitempty,oneof=y n"`
	RelaxedTypography              *string `validate:"omitempty,oneof=y u n"`
	DisambiguationApplied          *string `validate:"omitempty,oneof=n m s"`
	SemanticDisambiguationGrouping *string `validate:"omitempty,oneof=n g t l"`
	DisambiguationContext          *string `validate:"-"`
	UserDictionary                 *string `validate:"-"`
	TopicTypes                     *string `validate:"-"`
	ShowSubtopics                  *string `validate:"omitempty,oneof=y n"`
	TimeReference                  *string `validate:"-"`
	SentimentModel                 *string `validate:"omitempty,oneof=general"`
	ExpandGlobalPolarity           *string `validate:"omitempty,oneof=y n"`
}

func (p *Parser) Validate() error {
	validate := validator.New()

	if err := validate.Struct(p); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if p.Text == "" && p.URL == "" && p.Document == "" {
		return fmt.Errorf("one of the following fields must be set: txt, url, doc")
	}

	if p.Text != "" && (p.URL != "" || p.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if p.URL != "" && (p.Text != "" || p.Document != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	if p.Document != "" && (p.Text != "" || p.URL != "") {
		return fmt.Errorf("only one of the following fields can be set: txt, url, doc")
	}

	return nil

}

func (p *Parser) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := p.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", p.Key)
	multipartForm.AddField("lang", p.InputLanguage)
	multipartForm.AddOptionalField("of", p.OutputLanguage)
	multipartForm.AddMutualExclusiveFields(
		map[string]string{
			"txt": p.Text,
			"url": p.URL,
			"doc": p.Document,
		},
	)
	multipartForm.AddOptionalField("verbose", p.Verbose)
	multipartForm.AddOptionalField("txtf", p.TextFormat)
	multipartForm.AddOptionalField("uw", p.UnknownWords)
	multipartForm.AddOptionalField("rt", p.RelaxedTypography)
	multipartForm.AddOptionalField("dm", p.DisambiguationApplied)
	multipartForm.AddOptionalField("sdg", p.SemanticDisambiguationGrouping)
	multipartForm.AddOptionalField("cont", p.DisambiguationContext)
	multipartForm.AddOptionalField("ud", p.UserDictionary)
	multipartForm.AddOptionalField("tt", p.TopicTypes)
	multipartForm.AddOptionalField("st", p.ShowSubtopics)
	multipartForm.AddOptionalField("timeref", p.TimeReference)
	multipartForm.AddOptionalField("sm", p.SentimentModel)
	multipartForm.AddOptionalField("egp", p.ExpandGlobalPolarity)

	return multipartForm, nil
}

func (p *Parser) String() string {
	bytes, _ := json.MarshalIndent(p, "", " ")
	return string(bytes)
}
