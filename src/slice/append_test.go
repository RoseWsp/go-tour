package append_test

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {

	var s []int
	printSlice(s)

	/*探究 slice 的扩容*/
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2)
	printSlice(s) /* len=3 cap=4 [0 1 2] */

	s = append(s, 3, 4)
	printSlice(s)

	s = append(s, 5, 6, 7, 8)
	printSlice(s)

	var s1 []int

	s1 = append(s1, 0, 1, 2, 3, 4)
	printSlice(s1)

	s1 = append(s1, 4)
	printSlice(s1)

	s1 = append(s1, 6)
	printSlice(s1)

	s2 := []int{}

	s2 = append(s2, 1, 2, 3, 4)
	printSlice(s2)

	/*
		得出结论 何时扩容，扩容的大小是如何界定的
		扩容后的 cap 为扩容之前的1倍 ,即 new cap = 2 * old cap,如果len 超过了 new cap ,那么new cap = len + len % 2
	*/
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
