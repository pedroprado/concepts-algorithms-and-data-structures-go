package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	var arr = []int{1, 7, 5, 2}

	var expected = []int{1, 2, 5, 7}

	var received = bubbleSort(arr)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}
}

func TestBubbleSort_2(t *testing.T) {

	var arr = []int{1, 240, 2, 120, 7, 5, 1}

	var expected = []int{1, 1, 2, 5, 7, 120, 240}

	var received = bubbleSort2(arr)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}
}

//-------------------------Benchmark-----------------------
func BenchmarkBubbleSort(b *testing.B) {
	var arr = []int{1, 225, 235, 2, 332, 664, 3, 56, 34, 77, 225, 3, 5689, 9, 5986, 235, 4, 8, 9, 10, 19, 20}

	for i := 0; i < b.N; i++ {
		bubbleSort(arr)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	var arr = []int{1, 225, 235, 2, 332, 664, 3, 56, 34, 77, 225, 3, 5689, 9, 5986, 235, 4, 8, 9, 10, 19, 20}

	for i := 0; i < b.N; i++ {
		bubbleSort2(arr)
	}
}
