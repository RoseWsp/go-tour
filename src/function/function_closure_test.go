package function_test

import (
	"fmt"
	"testing"
)

func fibonacci() func() int {
	var (
		a = 1
		b = 1
	)
	call_times := 0
	return func() int {
		call_times++
		if call_times <= 2 {
			return 1
		} else {
			c := a + b
			a, b = b, c
			return c
		}
	}
}

func TestFunctionClosure(t *testing.T) {
	fib := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(fib())
	}
}
