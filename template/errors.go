package template

import "errors"

var (
	ErrInvalidConfig     = errors.New("Template: invalid Config object")
	ErrPoorlyWordedError = errors.New("Template: something went wrong")
)
