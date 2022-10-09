package response

import (
	"io"
)

type TextClassification struct {
	Response
	CategoryList []Category `json:"category_list"`
}

func (t *TextClassification) Decode(body io.ReadCloser) error {
	return decode(body, t)
}
