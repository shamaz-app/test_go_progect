package main

import "fmt"

func goSelectionSort() {
	data := getTrainData()

	fmt.Println("Before selection sort:")
	fmt.Println(data)

	selectionSort(&data)

	fmt.Println("After selection array:")
	fmt.Println(data)
}

func selectionSort(data *[]int) {
	for i := 0; i < len(*data); i ++ {
		minElemIndex := findMinElementFromIndex(data, i)
		swap(data, minElemIndex, i)
	}
}

func findMinElementFromIndex(data *[]int, fromIndex int) int {
	minElemIndex := fromIndex
	for i := fromIndex; i < len(*data); i++ {
		if (*data)[i] < (*data)[minElemIndex] {
			minElemIndex = i
		}
	}
	return minElemIndex
}

