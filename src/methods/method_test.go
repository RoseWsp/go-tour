package methods_test

import (
	"fmt"
	"math"
	"testing"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethods(t *testing.T) {
	v := Vertex{3, 5}
	v.Scale(10)
	fmt.Println(v.Abs())
}
