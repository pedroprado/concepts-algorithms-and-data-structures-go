package insertionsort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	arr := []int{1, 6, 3, 9, 2, 4}

	expected := []int{1, 2, 3, 4, 6, 9}
	received := insertionSort(arr)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)

	}

}

func TestMoveToBeggining(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}

	expected := []int{4, 1, 2, 3, 5, 6}
	received := moveToBeggining(arr, 3)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)

	}
}

func TestMoveToPlace(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}

	expected := []int{0, 1, 2, 6, 3, 4, 5}
	received := moveToPlace(arr, 3, 6)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}
}

//-----------------------------Benchmark--------------------------------
func BenchmarkSelectionSort_1000(b *testing.B) {

	var arr = generateCollection(1000)

	for i := 0; i < b.N; i++ {
		insertionSort(arr)
	}
}

func BenchmarkSelectionSort_10000(b *testing.B) {

	var arr = generateCollection(10000)

	for i := 0; i < b.N; i++ {
		insertionSort(arr)
	}
}

func generateCollection(size int) []int {

	var arr = []int{}

	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(size))
	}

	return arr
}
