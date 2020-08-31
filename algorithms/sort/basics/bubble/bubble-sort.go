package bubblesort

func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

	return arr
}

//This version is optimized for returning when it has not swapped any elements in a given iteration (thus,
//meaning it is already ordered)
//We also narrowed the range of the inner loop comparison ( j < len -1-i), because the last element is
//already the higher, thus it should not be compared
func bubbleSort2(arr []int) []int {

	var noSwaps bool

	for i := 0; i < len(arr)-1; i++ {
		noSwaps = true

		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				noSwaps = false
			}
		}

		if noSwaps {
			return arr
		}

	}

	return arr
}
