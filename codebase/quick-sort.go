package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	list := []int{22, 33, 5, 8, 32, 9, 28, -110, 6, 6, 23, 55, 54, 9}
	QuickSort(list, 0, len(list)-1)

	out, _ := json.Marshal(list)
	fmt.Printf("%s \n", string(out))
}

// start

func partition(list []int, i, j int) int {
	pivot := list[i]
	for i < j {
		if i < j && list[j] >= pivot {
			j--
		}
		list[i] = list[j]

		if i < j && list[i] <= pivot {
			i++
		}
		list[j] = list[i]
	}
	list[i] = pivot
	return i
}

func QuickSort(list []int, i, j int) {
	if j > i {
		pivotIndex := partition(list, i, j)

		// out, _ := json.Marshal(list)
		// fmt.Printf("pivot value: %d list: %s \n", list[pivotIndex], string(out))

		QuickSort(list, i, pivotIndex-1)
		QuickSort(list, pivotIndex+1, j)
	}
}
