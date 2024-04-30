package caller

import (
	"fmt"
)

func Error(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s, %w", NameSkip(2), err)
}
