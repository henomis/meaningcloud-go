package response

import (
	"encoding/json"
	"io"
)

type TextClassification struct {
	Response
	CategoryList []Category `json:"category_list"`
}

func (t *TextClassification) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(t)
}
