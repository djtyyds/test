package main

import "fmt"

func print(exitchan chan bool) {
	fmt.Println("下山的路又堵起了")
	exitchan <- true
	close(exitchan)
}

func main() {
	exitchan := make(chan bool, 1)
	go print(exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}
