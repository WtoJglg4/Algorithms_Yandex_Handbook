package main

import "fmt"

//to build: go build main.go AlgoFile.go
func main() {

	//HanoiTowers
	// s := 0
	// HanoiTowers(6, 1, 3, &s)
	// fmt.Println(s)
	// fmt.Println(FastHanoiTowers(6), FastHanoiTowers(4))

	//MergeSort
	// list := []int{7, 92, 87, 1, 4, 3, 2, 6}         //even number of elements
	// list_odd := []int{7, 92, 87, 1, 4, 3, 2, 6, 22} //odd number of elements
	// SortedList := MergeSort(list)
	// SortedList_odd := MergeSort(list_odd)
	// fmt.Println(SortedList, SortedList_odd)

	//QuickSort
	list := []int{7, 92, 87, 1, 4, 3, 2, 6}         //even number of elements
	list_odd := []int{7, 92, 87, 1, 4, 3, 2, 6, 22} //odd number of elements
	fmt.Println(list, list_odd)
	SortedList := QuickSort(list)
	SortedList_odd := QuickSort(list_odd)
	fmt.Println(SortedList, SortedList_odd)
}
