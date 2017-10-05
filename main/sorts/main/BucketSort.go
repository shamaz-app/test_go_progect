package main

import "fmt"

//func main() {
//	//Prepare data for sorting
//
//	data := []string{
//		"zipper",
//		"first_string",
//		"french",
//		"rube",
//		"rubi",
//		"ruby",
//		"rua",
//		"abc",
//		"eto",
//		"mayonez",
//		"21",
//		"12",
//		"41",
//		"1",
//		"ambassador",
//		"singleton",
//		"Jack!",
//		"Jagger",
//		"James"};
//
//
//	sort(&data)
//}


func sort(data *[]string) {
	fmt.Println("Start sorting")

	for _, item := range *data {
		fmt.Println(item)
	}
}