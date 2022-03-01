package main

import (
	"fmt"
	"math"
)

func merhaba(chan1 chan string) {
	chan1 <- "Merhaba"
}

type circle struct {
	r float64
}

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func main() {

	/*	myChannel := make(chan string)

		go merhaba(myChannel)

		fmt.Println(<-myChannel)*/
	c1 := circle{5}
	chan1 := make(chan float64)
	go area(c1, chan1)
	fmt.Printf("%2.f\n", <-chan1)
}
