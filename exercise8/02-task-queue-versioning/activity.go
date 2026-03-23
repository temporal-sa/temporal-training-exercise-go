package taskqueueversioning

import (
	"fmt"
	"strings"
)

func CreateGreeting(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func Uppercase(name string) (string, error) {
	return strings.ToUpper(name), nil
}
