package client

import "github.com/ricardoalcantara/api-email-client/pkg/types"

type ApiResponse[T any] struct {
	Status string               `json:"status"`
	Result T                    `json:"result"`
	Error  *types.ErrorResponse `json:"error"`
}
