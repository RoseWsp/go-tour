package methods_test

import (
	"fmt"
	"math"
	"testing"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func TestMyFloat(t *testing.T) {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
