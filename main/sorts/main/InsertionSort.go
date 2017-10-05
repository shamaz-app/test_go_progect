package main

import "fmt"

func goInsertionSort() {
	data := getTrainData()

	fmt.Println("Before insertion sort:")
	fmt.Println(data)

	insertionSort(&data)

	fmt.Println("After insertion sort:")
	fmt.Println(data)
}

func insertionSort(data *[]int) {
	for i := 0; i < len(*data); i ++ {
		for j := i; j > 0; j -- {
			if (*data)[j-1] > (*data)[j] {
				swap(data, j-1, j)
			} else {
				break
			}
		}
	}
}
