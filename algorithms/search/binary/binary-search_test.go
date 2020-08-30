package binarysearch

import (
	"testing"
)

func TestBinarySearch_success(t *testing.T) {
	var collection = []int{1, 2, 3, 4, 5}
	var searched = 4

	var expected = 3
	index, found := binarySearch(collection, searched)

	if !found {
		t.Errorf("\nExpected found: %v\nGot: %v\n", true, found)
	}

	if index != expected {
		t.Errorf("\nExpected index: %v\n Got: %v\n", expected, index)
	}
}

func TestBinarySearch_fail(t *testing.T) {
	var collection = []int{1, 2, 3, 4, 5}
	var searched = -1

	var expected = 0
	index, found := binarySearch(collection, searched)

	if found {
		t.Errorf("\nExpected found: %v\nGot: %v\n", false, found)
	}

	if index != expected {
		t.Errorf("\nExpected index: %v\n Got: %v\n", expected, index)
	}

}

//---------------------------BENCHING---------------------------
func BenchmarkBinarySearch_10000(b *testing.B) {
	var coll = generateCollection(10000)

	for i := 0; i < b.N; i++ {
		binarySearch(coll, 9999)
	}
}

func BenchmarkBinarySearch_100000(b *testing.B) {
	var coll = generateCollection(100000)

	for i := 0; i < b.N; i++ {
		binarySearch(coll, 99999)
	}
}

//----------------------------------------------------
func generateCollection(size int) []int {
	var collection = []int{}

	for i := 0; i < size; i++ {
		collection = append(collection, i)
	}

	return collection
}
