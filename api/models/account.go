package models

type CreateAccountRequest struct {
	Owner string `json:"owner"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
}

type UpdateAccountRequest struct {
	Owner *string `json:"owner"`
	Balance *int64 `json:"balance"`
	Currency *string `json:"currency"`
}