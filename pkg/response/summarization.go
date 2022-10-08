package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type Summarization struct {
	Response
	Summary string `json:"summary"`
}

func (s *Summarization) Ok() bool {
	return s.Status.Code == "0"
}

func (s *Summarization) Error() error {
	return fmt.Errorf("error code %s: %s", s.Status.Code, s.Status.Msg)
}

func (s *Summarization) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(s)
}
