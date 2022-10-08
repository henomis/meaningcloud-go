package response

import (
	"encoding/json"
	"fmt"
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

func (d *DocumentStructure) Ok() bool {
	return d.Status.Code == "0"
}

func (d *DocumentStructure) Error() error {
	return fmt.Errorf("error code %s: %s", d.Status.Code, d.Status.Msg)
}

func (d *DocumentStructure) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(d)
}
