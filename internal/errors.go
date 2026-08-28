package internal

import "fmt"

type HttpResponseError struct {
	Status     string
	StatusCode int
	Message    string
	Body       []byte
}

func (h *HttpResponseError) Error() string {
	return fmt.Sprintf("status: %s, statusCode: %v, message: %s, responseBody: %s", h.Status, h.StatusCode, h.Message, string(h.Body))
}
