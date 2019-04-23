package map_test

import (
	"fmt"
	"testing"
)

func TestMultating(t *testing.T) {

	m := make(map[string]int)

	m["Answer"] = 12
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
