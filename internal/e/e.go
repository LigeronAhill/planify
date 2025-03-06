package e

import (
	"fmt"
)

func Wrap(operation string, err error) error {
	return fmt.Errorf("операция: %s; ошибка: %s", operation, err.Error())
}
