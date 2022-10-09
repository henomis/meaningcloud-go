package response

import "fmt"

type Response struct {
	Status Status `json:"status"`
}

type Status struct {
	Code             string `json:"code"`
	Msg              string `json:"msg"`
	Credits          string `json:"credits"`
	RemainingCredits string `json:"remaining_credits"`
}

func (s *Status) IsSuccess() bool {
	return s.Code == "0"
}

func (s *Status) Error() error {
	return fmt.Errorf("error code %s: %s", s.Code, s.Msg)
}

// Uh!
type StatusIntValues struct {
	Code             int    `json:"code"`
	Msg              string `json:"msg"`
	Credits          int    `json:"credits"`
	RemainingCredits int    `json:"remaining_credits"`
}

func (s *StatusIntValues) IsSuccess() bool {
	return s.Code == 0
}

func (s *StatusIntValues) Error() error {
	return fmt.Errorf("error code %d: %s", s.Code, s.Msg)
}
