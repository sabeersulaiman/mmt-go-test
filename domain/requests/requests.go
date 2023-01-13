package requests

type AccountCreateRequest struct {
	Name           string  `json:"name"`
	IdNumber       string  `json:"idNumber"`
	InitialDeposit float64 `json:"initialDeposit"`
}

type AccountUpdateRequest struct {
	Amount float64 `json:"amount"`
}
