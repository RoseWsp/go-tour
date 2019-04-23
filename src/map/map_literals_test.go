package map_test

import (
	"fmt"
	"testing"
)

func TestMapLiterals(t *testing.T) {
	// m := map[string]Vertex{
	// 	"Bell Labs": Vertex{
	// 		40.5243535, -42.53345,
	// 	},
	// 	"Google": Vertex{
	// 		45.43535, -122.434355,
	// 	},
	// }
	// fmt.Println(m)

	m1 := map[string]Vertex{
		"Tencent":   {423.5335, -4253.5343},
		"Alilibaba": {425, 6446},
	}
	// fmt.Print(m1)

	for _, v := range m1 {
		fmt.Println(v.Lat)
	}
}
