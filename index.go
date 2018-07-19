package main

import "fmt"
import . "./allPackages"

func printaj() {
	//Da se napravat fukncii kako fmt.Println(funkcijaA())
	fmt.Println(Soberi(5, 10))
	fmt.Println(Mnozhi(5, 10))
	fmt.Println(Deli(100, 5))
	fmt.Println("example")
}

func main() {
	printaj()
}
