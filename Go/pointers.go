package main

import (
	"fmt"
	"math"
)

func main() {

	/*name := "arin"

	fmt.Println(name)
	fmt.Println(&name)*/

	/*	x := 22
		fmt.Println(x)
		fmt.Println(&x)
		fmt.Println(*(&x))*/

	/*x1 := 20
	x2 := &x1

	*x2 = 3
	fmt.Println(x1, *x2)*/

	/*a1 := [4]int{1, 2, 3, 4}
	a2 := a1

	a2[0] = 5
	fmt.Println(a2, a1)*/

	/*a1 := [4]int{1, 2, 3, 4}
	a2 := a1

	a2[0] = 5
	fmt.Println(a2, a1)
	*/

	/*	x := 5
		fmt.Println(x)
		double(x)
		fmt.Println(x)
	*/
	/*	x := 10
		y := &x
		*y = 20
		fmt.Println(x, *y)
	*/

	r1 := rectangle{5, 8}
	r1.display()
	fmt.Println(r1.area())
	fmt.Println(r1.circumference())

}

/*func double(num int) {
	num *= 2
	fmt.Println(num)
}*/
type rectangle struct {
	a, b int
}

func (r rectangle) display() {

	fmt.Println("Kenar Uzunlukları: ", r.a, "ve", r.b, "olan dikdörtgen")
}
func (r rectangle) area() int {
	return r.a * r.b
}
func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}





type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func (c circle) diameter() float64 {
	return 2 * c.r
}

type shape interface {
	area() float64
	circumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.circumference())
	fmt.Printf("%T", i)
	fmt.Println()
}


	/* 	r1 := rectangle{3, 8}
	   	fmt.Println("Area: ", r1.area())
	   	fmt.Println("Circumference: ", r1.circumference())
	   	fmt.Println()
		   interfaceFunc(r1) */


}

// concrete data type

r1 := rectangle{3, 8}
interfaceFunc(r1)

fmt.Println()
c1 := circle{5}
interfaceFunc(c1)


