package insertionsort

import (
	"github.com/google/uuid"
)

//PSEUDO CODE
//Starts a "growing list", with the first element of the collection
//Loop throught the rest of the collection
//For each element, find its place in the "growing list"
//Note: work with indexes

func insertionSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var subarrayIndex int = 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[subarrayIndex] {
			if arr[i] < arr[0] {
				arr = moveToBeggining(arr, i)

			} else {
				arr = findPlaceInSubarray(arr, subarrayIndex, i)
			}
		}

		subarrayIndex = subarrayIndex + 1
	}

	return arr

}

func moveToBeggining(arr []int, index int) []int {

	elemeToMove := arr[index]

	arr = append(arr[:index], arr[index+1:]...)

	arr = append([]int{elemeToMove}, arr...)

	return arr
}

func findPlaceInSubarray(arr []int, subindex int, elemIndex int) []int {
	var elem = arr[elemIndex]
	var indexToMove int = subindex
	for i := subindex; i >= 0; i-- {
		if elem < arr[i] {
			indexToMove = i
		}
	}
	arr = moveToPlace(arr, indexToMove, elemIndex)

	return arr
}

func moveToPlace(arr []int, indexToMove int, removeIndex int) []int {
	elemToMove := arr[removeIndex]

	arr = append(arr[:removeIndex], arr[removeIndex+1:]...)

	newarr := []int{elemToMove}
	newarr = append(newarr, arr[indexToMove:]...)

	arr = append(arr[:indexToMove], newarr...)

	return arr
}

func compareUUIDs(u1 uuid.UUID, u2 uuid.UUID) bool {

	return u1 == u2
}
