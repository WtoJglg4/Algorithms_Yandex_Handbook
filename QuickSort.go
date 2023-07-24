package main

func QuickSort(list []int) []int {
	n := len(list)
	if n == 1 {
		return list
	}

	var c_small, c_large, c_sorted []int
	for i, v := range list {
		if v < list[0] {
			c_small = append(c_small, list[i])
		}
		if v > list[0] {
			c_large = append(c_large, list[i])
		}
	}
	if len(c_small) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_small)...)
	}
	c_sorted = append(c_sorted, list[0])
	if len(c_large) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_large)...)
	}

	return c_sorted
}
