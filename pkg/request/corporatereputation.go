package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type CorporateReputation struct {
	Request
	InputLanguage     string  `validate:"required,oneof=auto en es fr it pt ca"`
	Text              string  `validate:"-"`
	Model             *string `validate:"omitempty,oneof=CorporateReputation"`
	Focus             *string `validate:"-"`
	Filter            *string `validate:"-"`
	RelaxedTypography *string `validate:"omitempty,oneof=y u n"`
}

func (c *CorporateReputation) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("invalid query %w", err)
	}

	if c.Text == "" {
		return fmt.Errorf("the following fields must be set: txt")
	}

	return nil

}

func (c *CorporateReputation) ToMultipartForm() (multipartform.MultipartForm, error) {

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	multipartForm := multipartform.New()

	multipartForm.AddField("key", c.Key)

	multipartForm.AddField("lang", c.InputLanguage)

	multipartForm.AddField("txt", c.Text)

	if c.Model != nil {
		multipartForm.AddField("model", *c.Model)
	}

	if c.Focus != nil {
		multipartForm.AddField("focus", *c.Focus)
	}

	if c.Filter != nil {
		multipartForm.AddField("filter", *c.Filter)
	}

	if c.RelaxedTypography != nil {
		multipartForm.AddField("rt", *c.RelaxedTypography)
	}

	return multipartForm, nil
}

func (c *CorporateReputation) String() string {
	bytes, _ := json.MarshalIndent(c, "", " ")
	return string(bytes)
}
