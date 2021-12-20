package functions

import (
	"fmt"
	"testing"
)

func TestDeferPanic(t *testing.T) {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
