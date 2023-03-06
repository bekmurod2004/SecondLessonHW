package fiz

import "fmt"

func Fizbuzer(a int) {
	// fizbuz

	for i := 1; i <= a; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, " fizzbuzz")
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println(i, " fizz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println(i, " buzz")
		} else {
			fmt.Println(i, "__")
		}
	}
	
}