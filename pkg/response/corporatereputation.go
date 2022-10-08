package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type CorporateReputation struct {
	Status     StatusIntValues `json:"status"` // Uh, this is not the standard status!
	Time       float64         `json:"time"`
	DeepTime   float64         `json:"deep_time"`
	EntityList []Entity        `json:"entity_list"`
}

func (c *CorporateReputation) Ok() bool {
	return c.Status.Code == 0
}

func (c *CorporateReputation) Error() error {
	return fmt.Errorf("error code %d: %s", c.Status.Code, c.Status.Msg)
}

func (c *CorporateReputation) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(c)
}