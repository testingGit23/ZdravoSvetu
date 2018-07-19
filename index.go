package main

import "fmt"
import "../ZdravoSvetu/pkg"
import "../ZdravoSvetu/pkg3"

func printaj() {
	//Da se napravat fukncii kako fmt.Println(funkcijaA())
	fmt.Println(soberiDvaBroja.SoberiDvaBroja(5, 10))
	//Gligor aj utre
	fmt.Println(dzalepackage.DeliDvaBroja(100, 5))
	fmt.Println("example")
}

func main() {
	printaj()
}
