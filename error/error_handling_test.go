package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func process(success bool) (string, error) {
	if success {
		return "done", nil
	}
	return "failed", &MyError{msg: "This is a custom error type"}
}
func TestCustomError(t *testing.T) {
	t.Run("error in golang is simple interface with a single method Error() string", func(t *testing.T) {
		result, err := process(false)
		assert.Equal(t, "This is a custom error type", err.Error())
		assert.Equal(t, "failed", result)
	})
}
