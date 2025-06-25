package apperror

import (
	"errors"
	"net/http"
	"strconv"
)

type Category int

const (
	ErrValidation Category = iota
	ErrInternal
	ErrNotFound
	ErrMethoNotAllowed
	ErrSecurity
	ErrForbidden
	ErrUnauthorized
)

// Code represents an error code with a category and an optional internal code.
// The Category field indicates the type of error, while the Internal field can be used
// to provide a specific error code for internal use.
type Code struct {
	Category Category
	Internal int
}

// AppError represents an application-specific error with additional metadata.
// It includes an error message, a status code, a category, and an optional internal code.
// The Metadata map can be used to store additional information about the error.
// The Error interface is implemented to allow easy error handling and logging.
type AppError struct {
	Err     error
	Status  int
	Code    Code
	Message string

	Metadata map[string]string
}

// NewAppError creates a new AppError with the provided error, category, and optional internal code.
// If internalCode is nil, it will not be included in the error code.
// The Status field can be used to set the HTTP status code associated with the error.
func NewAppError(err error, category Category, internalCode *int) *AppError {
	if internalCode != nil {
		return &AppError{
			Err: err,
			Code: Code{
				Category: category,
				Internal: *internalCode,
			},
			Metadata: make(map[string]string),
		}
	}

	return &AppError{
		Err: err,
		Code: Code{
			Category: category,
		},
		Metadata: make(map[string]string),
	}
}

// Error implements the error interface for AppError.
func (err AppError) Error() string {
	return err.Err.Error()
}

// BadRequest creates a new AppError with a status code of 400 (Bad Request).
func BadRequest(err error) *AppError {
	return withStatus(http.StatusBadRequest, err)
}

// NotFound creates a new AppError with a status code of 404 (Not Found).
func NotFound(err error) *AppError {
	return withStatus(http.StatusNotFound, err)
}

// Unauthorized creates a new AppError with a status code of 401 (Unauthorized).
func Unauthorized(err error) *AppError {
	return withStatus(http.StatusUnauthorized, err)
}

// Forbidden creates a new AppError with a status code of 403 (Forbidden).
func Forbidden(err error) *AppError {
	return withStatus(http.StatusForbidden, err)
}

// InternalServerError creates a new AppError with a status code of 500 (Internal Server Error).
func InternalServerError(err error) *AppError {
	return withStatus(http.StatusInternalServerError, err)
}

// WithField adds a field key value to the AppError's metadata.
func (err AppError) WithField(value string) *AppError {
	err.Metadata["field"] = value
	return &err
}

// WithRow adds a row key value number to the AppError's metadata.
func (err AppError) WithRow(row int) *AppError {
	err.Metadata["row"] = strconv.Itoa(row)
	return &err
}

// WithInfo adds additional information with "info" key to the AppError's metadata.
func (err AppError) WithInfo(info string) *AppError {
	err.Metadata["info"] = info
	return &err
}

// IsCategory checks if the provided error belongs to the specified category.
func IsCategory(srcErr error, category Category) bool {
	var appErr *AppError
	if errors.As(srcErr, &appErr) {
		return appErr.Code.Category == category
	}

	return false
}

func (c Category) String() string {
	switch c {
	case ErrValidation:
		return "ValidationError"
	case ErrInternal:
		return "InternalError"
	case ErrNotFound:
		return "NotFouncError"
	case ErrMethoNotAllowed:
		return "MethoNotAllowedError"
	case ErrForbidden:
		return "ForbiddenError"
	case ErrUnauthorized:
		return "UnauthorizedError"
	default:
		return "UnkownCategoryError"
	}
}

func (err AppError) Unwrap() error {
	return err.Err
}

func withStatus(internalCode int, err error) *AppError {
	return &AppError{
		Err: err,
		Code: Code{
			Internal: internalCode,
		},
		Message: err.Error(),
	}
}
