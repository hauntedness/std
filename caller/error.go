package caller

import (
	"fmt"
)

func Error(err error) error {
	return fmt.Errorf("%s, %w", NameSkip(2), err)
}
