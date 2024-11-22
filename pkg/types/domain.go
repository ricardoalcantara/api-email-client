package types

type ListView[T any] struct {
	List []T `json:"list"`
	Page int `json:"page"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
