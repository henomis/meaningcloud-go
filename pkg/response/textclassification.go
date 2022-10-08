package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type TextClassification struct {
	Response
	CategoryList []Category `json:"category_list"`
}

func (t *TextClassification) Ok() bool {
	return t.Status.Code == "0"
}

func (t *TextClassification) Error() error {
	return fmt.Errorf("error code %s: %s", t.Status.Code, t.Status.Msg)
}

func (t *TextClassification) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(t)
}
