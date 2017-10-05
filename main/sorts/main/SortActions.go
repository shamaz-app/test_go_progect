package main

func main() {
	goSelectionSort()
	goInsertionSort()
	goMergeSort()
}

func swap(data *[]int, i int, j int) {
	if i != j {
		temp := (*data)[i]
		(*data)[i] = (*data)[j]
		(*data)[j] = temp
	}
}

func getTrainData() []int {
	return []int{10, 8, 7, 1, 3, 12, 5, 4, 2}
}