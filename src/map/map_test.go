package map_test

import (
	"fmt"
	"testing"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func TestMap(t *testing.T) {

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		124.42, 632.6,
	}

	fmt.Println(m["Bell Labs"])

}
