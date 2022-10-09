package response

import (
	"encoding/json"
	"io"
)

type Summarization struct {
	Response
	Summary string `json:"summary"`
}

func (s *Summarization) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(s)
}
