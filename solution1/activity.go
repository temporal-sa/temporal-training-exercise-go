package solution1

import "fmt"

func CreateGreeting(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
