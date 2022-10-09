package response

import (
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

func (d *DeepCategorization) Decode(body io.ReadCloser) error {
	return decode(body, d)
}
