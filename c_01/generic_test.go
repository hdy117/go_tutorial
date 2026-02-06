package main

import (
	"fmt"
	"testing"
)

func TestAddTemplate(t *testing.T) {
	var a int = 5
	var b int = 6
	c := AddTemplate[int](a, b)

	fmt.Printf("result is %d.\n", c)
}

func TestAddTempalte1(t *testing.T) {
	var a int = 5
	var b float32 = 6.0
	c := AddTemplate1(a, b)
	fmt.Printf("c is %f.\n", c)
}

func TestQueue(t *testing.T) {
	var q QueueFloat64Ptr = &QueueFloat64{elements: make([]*float64, 0, 100)}
	a := 1.0
	b := 2.0
	c := 3.0
	d := 9.0
	q.Push(&a)
	q.Push(&b)
	q.Push(&c)
	q.Push(&d)

	for i := range 4 {
		ele, ok := q.Pop()
		fmt.Printf("idx:%d, pop val:%f, %t.\n", i, *ele, ok)
	}
}
