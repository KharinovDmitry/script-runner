package dto

import (
	"database/sql"
	"errors"
	"net/http"
)

type ApiError struct {
	StatusCode int
	Message    string
}

func NewApiError(err error) ApiError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ApiError{StatusCode: http.StatusNotFound}
	default:
		return ApiError{StatusCode: http.StatusInternalServerError}
	}
}
