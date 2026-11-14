package models

type CreateRequest struct {
	Fields map[string]string
	Photos []string
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
