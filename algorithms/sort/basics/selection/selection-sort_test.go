package selectionsort

import (
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {

	// go func() {
	// 	<-time.After(time.Second * 3)
	// 	panic("Timed out! Tests took too long!")
	// }()

	os.Exit(m.Run())
}

func TestSelectionSort(t *testing.T) {

	var arr = []int{1, 9, 28, 2, 7}

	var expected = []int{1, 2, 7, 9, 28}

	var received = selectionSort(arr)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, received)
	}
}

//-----------------------------Benchmark--------------------------------
func BenchmarkSelectionSort_1000(b *testing.B) {

	var arr = generateCollection(1000)

	for i := 0; i < b.N; i++ {
		selectionSort(arr)
	}
}

func BenchmarkSelectionSort_10000(b *testing.B) {

	var arr = generateCollection(10000)

	for i := 0; i < b.N; i++ {
		selectionSort(arr)
	}
}

func generateCollection(size int) []int {

	var arr = []int{}

	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(size))
	}

	return arr
}
