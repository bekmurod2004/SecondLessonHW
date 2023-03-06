package fib

import (
	"fmt"

)

func Fib(a int) {

	// fibonacci
	
	n1 := 0
	n2 := 1
	ans := 0
	
	sum:=0
	
	// 0 1 1 2 3 5 8 13 21 34 55

	for i := 2; i <= a; i++ {
		sum = n1 + n2
		ans += sum
		
		


		n1 = n2
		n2 = sum
	}
	fmt.Println(ans +1)
	
}