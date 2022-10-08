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

	if p.OutputLanguage != nil {
		multipartForm.AddField("of", *p.OutputLanguage)
	}

	if p.Text != "" {
		multipartForm.AddField("txt", p.Text)
	} else if p.URL != "" {
		multipartForm.AddField("url", p.URL)
	} else if p.Document != "" {
		multipartForm.AddFile("doc", p.Document)
	}

	if p.Verbose != nil {
		multipartForm.AddField("verbose", *p.Verbose)
	}

	if p.TextFormat != nil {
		multipartForm.AddField("txtf", *p.TextFormat)
	}

	if p.UnknownWords != nil {
		multipartForm.AddField("uw", *p.UnknownWords)
	}

	if p.RelaxedTypography != nil {
		multipartForm.AddField("rt", *p.RelaxedTypography)
	}

	if p.DisambiguationApplied != nil {
		multipartForm.AddField("dm", *p.DisambiguationApplied)
	}

	if p.SemanticDisambiguationGrouping != nil {
		multipartForm.AddField("sdg", *p.SemanticDisambiguationGrouping)
	}

	if p.DisambiguationContext != nil {
		multipartForm.AddField("cont", *p.DisambiguationContext)
	}

	if p.UserDictionary != nil {
		multipartForm.AddField("ud", *p.UserDictionary)
	}

	if p.TopicTypes != nil {
		multipartForm.AddField("tt", *p.TopicTypes)
	}

	if p.ShowSubtopics != nil {
		multipartForm.AddField("st", *p.ShowSubtopics)
	}

	if p.TimeReference != nil {
		multipartForm.AddField("timeref", *p.TimeReference)
	}

	if p.SentimentModel != nil {
		multipartForm.AddField("sm", *p.SentimentModel)
	}

	if p.ExpandGlobalPolarity != nil {
		multipartForm.AddField("egp", *p.ExpandGlobalPolarity)
	}

	return multipartForm, nil
}

func (p *Parser) String() string {
	bytes, _ := json.MarshalIndent(p, "", " ")
	return string(bytes)
}
