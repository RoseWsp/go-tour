package function_test

import (
	"fmt"
	"math"
	"testing"
)

func compute(x, y float64, fn func(float64, float64) float64) float64 {
	return fn(x, y)
}
func TestFunctionValue(t *testing.T) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(3, 4, hypot))
	fmt.Println(compute(3, 4, math.Pow))
}
