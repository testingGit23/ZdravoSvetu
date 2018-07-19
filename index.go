package main

import "fmt"
import "../ZdravoSvetu/pkg"
import "../ZdravoSvetu/pkg2"

func printaj() {
	//Da se napravat fukncii kako fmt.Println(funkcijaA())
	fmt.Println(soberiDvaBroja.SoberiDvaBroja(5, 10))
	//Gligor aj utre
	fmt.Println(mnozi.MnoziDvaBroja(5, 10))
	//Darko
	fmt.Println("example")
}

func main() {
	printaj()
}
