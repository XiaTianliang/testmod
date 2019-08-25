package testmod

import "fmt"

func Hi(name, text string) (string, error) {
	return fmt.Sprintf("Hi, %s %s!", name, text), nil
}
