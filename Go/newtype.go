package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	m1 := mile(10)
	k1 := toKilometer(m1)
	fmt.Println(k1)

	k2 := kilometer(20)
	m2 := toMile(k2)
	fmt.Println(m2)

}
func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)

}
func toMile(k kilometer) mile {

	return mile(k * 0.62)
}
