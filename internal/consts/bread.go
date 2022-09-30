package consts

import "fmt"

var (
	CannotFoundIdError   = newCannotFoundError("id", 0)
	CannotFoundNameError = newCannotFoundError("name", "")
)

type CannotFoundError struct {
	name string
	v    interface{}
}

func newCannotFoundError(name string, v interface{}) *CannotFoundError {
	return &CannotFoundError{name: name, v: v}
}

func (e *CannotFoundError) Name(name string) *CannotFoundError {
	c := newCannotFoundError(name, e.v)
	return c
}

func (e *CannotFoundError) Value(value interface{}) *CannotFoundError {
	c := newCannotFoundError(e.name, value)
	return c
}

func (e *CannotFoundError) Error() string {
	return fmt.Sprintf("Cannot find Hello with %s %q", e.name, e.v)
}
