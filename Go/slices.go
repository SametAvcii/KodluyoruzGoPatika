package main

import (
	"fmt"
)

type employee struct {
	name      string
	age       int
	isMarried bool
	kids      []string
}
type manager struct {
	employee
	hasDegree bool
}

// Klasik OOP'de Is A kuralı vardır. Yani Aslan hayvan sınıfında olduğu durumunda aynı zamanda bir hayvandır.
//Hayvanların tüm özellikleriini taşırAynı zamanda kendşne has özellikleride olabilir.

//Fakat Go'da Has A kuralı vardır. Yani burada tanımladığımız manager aynı zamanda bir employee değildir.

func main() {

	var e1 employee
	e1.name = "SAmet"
	e1.age = 20
	e1.isMarried = false
	e1.kids = []string{
		"Yusuf",
		"Ahmet",
	}

	var e2 employee
	e2.name = "Aykut"
	e2.age = 20
	e2.isMarried = true

	fmt.Println(e1)
	fmt.Println(e2)

	e3 := employee{
		name:      "Mete",
		age:       21,
		isMarried: false,
	}
	fmt.Println(e3)

	m1 := manager{
		employee: employee{
			name:      "Anıl",
			age:       20,
			isMarried: false,
		},
		hasDegree: true,
	}
	fmt.Println(m1)

	theBoss := struct {
		name  string
		money bool
	}{
		name:  "theBoss",
		money: true,
	}
	fmt.Println(theBoss)
}
