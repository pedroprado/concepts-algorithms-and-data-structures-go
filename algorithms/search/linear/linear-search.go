package linearsearch

func linearSearch(nums []int, searched int) *int {

	for i, elem := range nums {
		if elem == searched {
			return &i
		}
	}

	return nil
}
