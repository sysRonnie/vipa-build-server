package expense

import "fmt"

type DuplicateClientNameError struct {
	Name string
}

func (e *DuplicateClientNameError) Error() string {
	return fmt.Sprintf("Client '%s' already exists", e.Name)
}