package mergesort

// IDEA:
// Divide the collection to N ollections of 1 element, spliting it in half for each iteration
// Reorder and merge them, two by two, till one have the original collection
// Use recursion for spliting the collection

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := []int{}
	right := []int{}

	for i := 0; i < len(arr)/2; i++ {
		left = append(left, arr[i])
	}
	for i := len(arr) / 2; i < len(arr); i++ {
		right = append(right, arr[i])
	}

	left = mergeSort(left)
	right = mergeSort(right)

	results := merge(left, right)

	return results
}

//Compare the first element of both arrays, and put the lower on in the results array
func merge(left []int, right []int) []int {

	results := []int{}

	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}

		if left[0] <= right[0] {
			results = append(results, left[0])

			//remove the left[0] from left
			left = append([]int{}, left[1:]...)
		} else {
			results = append(results, right[0])
			//remove the right[0] from right
			right = append([]int{}, right[1:]...)

		}
	}

	//put the rest of the right arr
	if len(left) == 0 && len(right) != 0 {
		for _, v := range right {
			results = append(results, v)
		}
	}

	//put the rest of the left arr
	if len(left) != 0 && len(right) == 0 {
		for _, v := range left {
			results = append(results, v)
		}
	}

	return results
}
