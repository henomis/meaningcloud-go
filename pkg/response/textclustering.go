package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type TextClustering struct {
	Response
	ClusterList []Cluster `json:"cluster_list"`
}

type Cluster struct {
	Title        string            `json:"title"`
	Size         int               `json:"size"` // Uh, this is not a string as in the docs
	Score        string            `json:"score"`
	DocumentList map[string]string `json:"document_list"` // Uh, this is not a list!
}

func (t *TextClustering) Ok() bool {
	return t.Status.Code == "0"
}

func (t *TextClustering) Error() error {
	return fmt.Errorf("error code %s: %s", t.Status.Code, t.Status.Msg)
}

func (t *TextClustering) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(t)
}
