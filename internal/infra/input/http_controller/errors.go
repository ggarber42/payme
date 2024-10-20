package http_controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	HTTPStatusCode int    `json:"status"`
	StatusText     string `json:"status_text"`
	ErrorText      string `json:"error"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func NewErrorResponse(httpStatus int, statusText string, err error) *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: httpStatus,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}
