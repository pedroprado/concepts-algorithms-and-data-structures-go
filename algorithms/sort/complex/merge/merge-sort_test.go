package mergesort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{29, 7, 11, 9, 20, 13}

	expected := []int{7, 9, 11, 13, 20, 29}
	received := mergeSort(arr)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}

}

func TestMerge(t *testing.T) {

	left := []int{7, 13, 20}
	right := []int{9, 11, 29}

	expected := []int{7, 9, 11, 13, 20, 29}

	received := merge(left, right)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}

}

//------------------Benchmark-----------------------------
func BenchmarkMergeSort_1000(b *testing.B) {
	arr := generateCollection(1000)

	for i := 0; i < b.N; i++ {
		mergeSort(arr)
	}
}

func BenchmarkMergeSort_10000(b *testing.B) {
	arr := generateCollection(10000)

	for i := 0; i < b.N; i++ {
		mergeSort(arr)
	}
}

// func BenchmarkMergeSort_100000(b *testing.B) {
// 	arr := generateCollection(100000)

// 	for i := 0; i < b.N; i++ {
// 		mergeSort(arr)
// 	}
// }

func generateCollection(size int) []int {

	var arr = []int{}

	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(size))
	}

	return arr
}
