package selectionsort

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		var lowest int = i

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[lowest] {
				lowest = j
			}

		}

		arr[i], arr[lowest] = arr[lowest], arr[i]
	}

	return arr
}
