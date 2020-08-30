package binarysearch

// searchs in a ORDERED collection
// PSEUDO CODE
// Split the collection in two (left and right), and check in which of them is the searched number

func binarySearch(nums []int, searched int) (int, bool) {

	var left = 0
	var right = len(nums) - 1
	var middle int

	for left != right {
		middle = (left + right) / 2

		if searched > nums[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if nums[left] == searched {
		return left, true
	}

	return 0, false
}
