package main

import (
	"fmt"
	"getcelcius"
)

func main() {
	fmt.Print("lütfen sıcaklığı giriniz:")
	celcius, err := GetCelcius.GetCelcius()
	if err != nil {
		fmt.Println(err)
	}
	kelvin := celcius + 273
	fmt.Printf("girdiğiniz değer %v celcius ve %v kelvindir", celcius, kelvin)

}
