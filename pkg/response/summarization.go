package response

import (
	"io"
)

type Summarization struct {
	Response
	Summary string `json:"summary"`
}

func (s *Summarization) Decode(body io.ReadCloser) error {
	return decode(body, s)
}
