package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func RandomizedQuickSort(list []int) []int {
	n := len(list)
	if n == 1 {
		return list
	}

	// Инициализация генератора случайных чисел с использованием текущего времени в качестве источника
	rand.NewSource(time.Now().UnixNano())
	m := rand.Intn(n - 1)
	fmt.Println(m)
	var c_small, c_large, c_sorted []int
	for i, v := range list {
		if v < list[m] {
			c_small = append(c_small, list[i])
		}
		if v > list[m] {
			c_large = append(c_large, list[i])
		}
	}
	if len(c_small) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_small)...)
	}
	c_sorted = append(c_sorted, list[m])
	if len(c_large) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_large)...)
	}

	return c_sorted
}
