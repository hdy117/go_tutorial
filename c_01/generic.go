package main

func AddTemplate[T int | int32 | int16 | int8 | float32](a T, b T) T {
	return a + b
}

func AddTemplate1[T1 int | int32 | int16 | int8 | float32, T2 float32](a T1, b T2) T2 {
	return T2(a) + b
}

// Queue
type Queue[T any] struct {
	elements []*T // slice of elements pointers
}

type QueueFloat64 = Queue[float64]
type QueueFloat64Ptr = *Queue[float64]

// Push element pointer into the end of queue
func (obj *Queue[T]) Push(element *T) {
	obj.elements = append(obj.elements, element)
}

// Pop the front pointer of elements from queue
func (obj *Queue[T]) Pop() (*T, bool) {
	sizeElements := len(obj.elements)
	if sizeElements == 0 {
		return nil, false
	}

	// pop the first
	element := obj.elements[0]
	obj.elements = obj.elements[1:sizeElements]

	return element, true
}
