package response

import (
	"encoding/json"
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

func (t *TextClustering) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(t)
}
