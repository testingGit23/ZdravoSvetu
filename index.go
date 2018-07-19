package main

import "fmt"
import . "./allPackages"

func main() {
	//Da se napravat fukncii kako fmt.Println(funkcijaA())
	fmt.Println(Soberi(5, 10))
	fmt.Println(Mnozhi(5, 10))
	fmt.Println(Deli(100, 5))

	twoD := SoberiNizaa(2, 3)
	for i := 0; i < 2; i++ {
		fmt.Println(twoD[i])
	}

	fmt.Println(SoberiNizaa(2, 3))
	fmt.Println(SoberiFloati(2.3, 5.4))
	fmt.Println("example")
}
