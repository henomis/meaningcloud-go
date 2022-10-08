package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type DeepCategorization struct {
	Response
	CategoryList []Category `json:"category_list,omitempty"`
}

type Category struct {
	Code              string `json:"code"`
	Label             string `json:"label"`
	AbsoluteRelevance string `json:"abs_relevance"`
	Relevance         string `json:"relevance"`
	Popolarity        string `json:"popolarity"`
	TermList          []Term `json:"term_list,omitempty"`
}

type Term struct {
	AbsoluteRelevance string   `json:"abs_relevance"`
	Form              string   `json:"form"`
	OffsetList        []Offset `json:"offset_list,omitempty"`
}

type Offset struct {
	Endp string `json:"endp"`
	Inip string `json:"inip"`
}

func (d *DeepCategorization) Ok() bool {
	return d.Status.Code == "0"
}

func (d *DeepCategorization) Error() error {
	return fmt.Errorf("error code %s: %s", d.Status.Code, d.Status.Msg)
}

func (d *DeepCategorization) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(d)
}
