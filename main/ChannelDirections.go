package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 3)
	pongs := make(chan string, 3)
	ping(pings, "passed message")
	ping(pings, "passed message(2)")
	ping(pings, "passed message(3)")
	pong(pings, pongs)
	pong(pings, pongs)
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
}