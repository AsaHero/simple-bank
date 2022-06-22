package models

type CreateAccountRequest struct {
	Owner string `json:"owmer"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
}
