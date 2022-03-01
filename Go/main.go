package main

import "fmt"

func main() {
	//cities := [4]string{"maras", "antep", "istanbul"}
	cities := []string{"maras", "antep", "istanbul"}

	fmt.Println(cities)
	fmt.Println(len(cities))
	//uznuluğu verir.

	var myArr [5]int
	fmt.Println(myArr)
	//[0 0 0 0 0] çıktı bu şekilde olur. Default değer O'dır integer tipi için.
	myArr[1] = 10
	fmt.Println(myArr[1])
	//Bu şekilde değer atanabilir.

	myArr1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myArr1 = mySquare(myArr1)
	fmt.Println(myArr1)
}
func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
