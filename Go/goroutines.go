package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go printX()
	wg.Wait()
	fmt.Println()
	printY()

}
func printX() {

	for i := 0; i < 20; i++ {
		fmt.Print("X")

	}
	wg.Done()
}
func printY() {

	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}
