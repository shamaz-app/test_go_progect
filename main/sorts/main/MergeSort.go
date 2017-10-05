package main

import "fmt"

type mergePart struct {
	part []int
}

func goMergeSort() {
	data := getTrainData()

	fmt.Println("Before merge sort:")
	fmt.Println(data)

	mergeSort(&data)

	fmt.Println("After merge sort:")
	fmt.Println(data)
}

func mergeSort(data *[]int) {
	if len(*data) <= 1 {
		return
	}
	sliceData := make([]int, len(*data))
	copy(sliceData, *data)

	firstPiece := sliceData[:len(*data) / 2]
	secondPiece := sliceData[len(*data) / 2:]

	mergeSort(&firstPiece)
	mergeSort(&secondPiece)

	merge(data, firstPiece, secondPiece)
}

func merge(data *[]int, left []int, right []int) {
	leftIndex := 0;
	rightIndex := 0;
	targetIndex := 0;
	remaining := len(left) + len(right);
	for remaining > 0 {
		if leftIndex >= len(left) {
			(*data)[targetIndex] = right[rightIndex]
			rightIndex ++;
		} else if rightIndex >= len(right) {

			(*data)[targetIndex] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] <= right[rightIndex] {
			(*data)[targetIndex] = left[leftIndex]
			leftIndex++
		} else {
			(*data)[targetIndex] = right[rightIndex]
			rightIndex ++;
		}

		targetIndex++;
		remaining--;
	}
}

//we
//func split(data []int) []mergePart {
//	temp := make([]mergePart, len(data))
//
//	sliceData := make([]int, len(data))
//	copy(sliceData, data)
//
//	firstPiece := sliceData[:len(data)]
//	secondPiece := sliceData[len(data):]
//
//	for item := range split(firstPiece) {
//		temp = append(temp, item.part)
//	}
//	temp = append(temp, splitResult)
//
//	temp = append(temp, , split(secondPiece))
//
//	return temp
//}