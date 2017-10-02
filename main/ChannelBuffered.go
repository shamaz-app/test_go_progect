package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "ping"
	messages <- "ping2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}