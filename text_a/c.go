package main

import (
	"fmt"
	"time"
)

func print1() {
	fmt.Println("你好")

}

func print2() {
	fmt.Println("我不好")
}

func print3() {
	fmt.Println("希望考试不挂")
}

func print4() {
	fmt.Println("4级")
}

func print5() {
	fmt.Println("6级")
}

func print6() {
	fmt.Println("计算机2级")
}

func print7() {
	fmt.Println("证书")
}

func print8() {
	fmt.Println("23")
}

func print9() {
	fmt.Println("嗯")
}

func print10(exitchan chan bool) {
	fmt.Println("慢慢来")
	time.Sleep(time.Second * 2)
	exitchan <- true
	close(exitchan)
}

func main() {
	exitchan := make(chan bool, 1)
	go print1()
	go print2()
	go print4()
	go print3()
	go print5()
	go print6()
	go print7()
	go print8()
	go print9()
	go print10(exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}
