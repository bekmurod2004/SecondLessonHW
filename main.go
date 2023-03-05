package main

import "fmt"

func main() {
	var number int 
	fmt.Scanln(&number)
	fmt.Println()

	fizbuzer(number)
	fmt.Println()

	fib(number)
	fmt.Println()
	
	fact(number)
}