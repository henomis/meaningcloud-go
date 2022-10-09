package request

import (
	"fmt"
	"strings"
)

type Request struct {
	Key string `validate:"required"`
}

func validateMutualExclusiveFields(fields map[string]string) error {

	var keys []string
	emptyFields := 0
	nonEmptyFields := 0

	for key, value := range fields {
		keys = append(keys, key)
		if value == "" {
			emptyFields++
		} else {
			nonEmptyFields++
		}
	}

	if emptyFields == len(fields) {
		return fmt.Errorf("one of the following fields must be set: %s", strings.Join(keys, ", "))
	}

	if nonEmptyFields > 1 {
		return fmt.Errorf("only one of the following fields can be set: %s", strings.Join(keys, ", "))
	}

	return nil
}
