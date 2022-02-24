package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*x, y := 10, 4
	sum, dif, prod := calc(x, y)
	fmt.Println(sum)
	fmt.Println(dif)
	fmt.Println(prod) */

	/*fmt.Print("Lütfen notunuzu giriniz:")
	grade, err := getGrade()
	if err != nil {
		fmt.Println(err)
	}
	if grade < 50 {
		fmt.Println("Kaldınız")
	} else {
		fmt.Println("Geçtiniz")
	}*/

}
func calc(num1, num2 int) (int, int, int) {
	sum := num1 + num2
	dif := num1 - num2
	prod := num1 * num2
	return sum, dif, prod
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}