package main

func MergeSort(list []int) []int {

	n := len(list)
	if n == 1 {
		return list
	}
	var FirstHalf, SecondHalf []int
	if n%2 == 0 {
		FirstHalf = list[:n/2]
		SecondHalf = list[n/2:]

	} else {
		FirstHalf = list[:n/2+1]
		SecondHalf = list[n/2+1:]
	}

	SortedFH := MergeSort(FirstHalf)
	SortedSH := MergeSort(SecondHalf)

	SortedList := Merge(SortedFH, SortedSH)
	return SortedList
}

func Merge(list_1, list_2 []int) []int {
	var SortedList []int

	for len(list_1) > 0 || len(list_2) > 0 {
		if len(list_1) == 0 {
			SortedList = append(SortedList, list_2...)
			return SortedList
		}
		if len(list_2) == 0 {
			SortedList = append(SortedList, list_1...)
			return SortedList
		}
		if list_1[0] < list_2[0] {
			SortedList = append(SortedList, list_1[0])
			list_1 = list_1[1:]
		} else {
			SortedList = append(SortedList, list_2[0])
			list_2 = list_2[1:]
		}
	}

	return SortedList
}
