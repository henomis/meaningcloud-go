package response

type Response struct {
	Status Status `json:"status"`
}

type Status struct {
	Code             string `json:"code"`
	Msg              string `json:"msg"`
	Credits          string `json:"credits"`
	RemainingCredits string `json:"remaining_credits"`
}

// Uh!
type StatusIntValues struct {
	Code             int    `json:"code"`
	Msg              string `json:"msg"`
	Credits          int    `json:"credits"`
	RemainingCredits int    `json:"remaining_credits"`
}
