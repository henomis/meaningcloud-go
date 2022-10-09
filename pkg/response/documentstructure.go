package response

import (
	"io"
)

type DocumentStructure struct {
	Response
	Title        string     `json:"title"`
	HeadingList  []string   `json:"heading_list"`
	AbstractList []string   `json:"abstract_list"`
	EmailsInfo   EmailsInfo `json:"emails_info"`
}

type EmailsInfo struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Subject []string `json:"subject"`
}

func (d *DocumentStructure) Decode(body io.ReadCloser) error {
	return decode(body, d)
}
