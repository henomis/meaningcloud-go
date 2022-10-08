package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type LanguageIdentification struct {
	Status       StatusIntValues `json:"status"` // Uh, this is not a standard status!
	LanguageList []Language      `json:"language_list"`
}

type Language struct {
	Language  string  `json:"language"`
	Name      string  `json:"name"`
	Relevance float64 `json:"relevance"`
	ISO6391   string  `json:"iso-639-1"`
	ISO6392   string  `json:"iso-639-2"`
	ISO6393   string  `json:"iso-639-3"`
	ISO6395   string  `json:"iso-639-5"`
	DeepTime  string  `json:"deep_time"`
	Time      string  `json:"time"`
}

func (l *LanguageIdentification) Ok() bool {
	return l.Status.Code == 0
}

func (l *LanguageIdentification) Error() error {
	return fmt.Errorf("error code %d: %s", l.Status.Code, l.Status.Msg)
}

func (l *LanguageIdentification) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(l)
}
