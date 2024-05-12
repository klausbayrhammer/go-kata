package main

import "fmt"

func DoThings(resultChan chan string) {
	// things are done here...
	// and the result sent via the channel
	resultChan <- "result of a loooong computation"
}

func main() {
	ch := make(chan string)

	go DoThings(ch) // run concurrently

	res := <-ch // block execution waiting to receive from channel
	fmt.Println(res)
}
