package linearsearch

import (
	"reflect"
	"testing"
)

func TestLinearSeach_success(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}

	var searched = 5

	var expected = 4
	var result = linearSearch(nums, searched)

	if !reflect.DeepEqual(expected, *result) {
		t.Errorf("\nExpected : %v\nGot: %v", expected, *result)
	}
}

func TestLinearSeach_failed(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}

	var searched = 6

	var result = linearSearch(nums, searched)

	if result != nil {
		t.Errorf("\nExpected : %v\nGot: %v", nil, result)
	}
}

//------------------Benching-------------------
func BenchmarkLinearSearch_10000(b *testing.B) {

	coll := generateCollection(10000)

	for i := 0; i < b.N; i++ {
		linearSearch(coll, 9999)
	}

}

func BenchmarkLinearSearch_100000(b *testing.B) {

	coll := generateCollection(100000)

	for i := 0; i < b.N; i++ {
		linearSearch(coll, 99999)
	}
}

//-----------------------------------------------------------
func generateCollection(size int) []int {
	var collection = []int{}

	for i := 0; i < size; i++ {
		collection = append(collection, i)
	}

	return collection
}
