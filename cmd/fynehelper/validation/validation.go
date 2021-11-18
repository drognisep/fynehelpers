package validation

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrValidation = errors.New("validation error")
)

func fieldValidationError(msg string, fields ...string) error {
	if msg == "" {
		msg = "Fields failed validation"
	}
	return errors.Wrap(ErrValidation, fmt.Sprintf("%s: %s", msg, strings.Join(fields, ", ")))
}

func NoneBlank(fields map[string]string) error {
	var blankFields []string
	for field, val := range fields {
		if val == "" {
			blankFields = append(blankFields, field)
		}
	}
	if len(blankFields) > 0 {
		return fieldValidationError("field(s) are blank", blankFields...)
	}
	return nil
}
