package testmod

import "fmt"

func Hi(name string) (string, error) {
	return fmt.Sprintf("Hi %s", name), nil
}
