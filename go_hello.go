package go_hello

import (
	"fmt"
)

func Show(message string) string{
	return fmt.Sprintf("First module - Hello, %s", message)
}