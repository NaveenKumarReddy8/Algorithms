package sorting

func performInsertionSort(toBeSorted []int) []int {
	lengthOfSlice := len(toBeSorted)
	if lengthOfSlice == 1 {
		return toBeSorted
	}

	for index := 1; index < lengthOfSlice; index++ {
		key := toBeSorted[index]
		previousIndex := index - 1
		for ; previousIndex >= 0; previousIndex-- {
			if toBeSorted[previousIndex] > key {
				toBeSorted[previousIndex+1] = toBeSorted[previousIndex]
			} else {
				break
			}
		}
		toBeSorted[previousIndex+1] = key
	}
	return toBeSorted
}

// Insertion .
func Insertion(integerInput []int) []int {
	return performInsertionSort(integerInput)
}
