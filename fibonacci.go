package main

import "fmt"

func fib(a int) {

	// fibonacci
	
	n1 := 0
	n2 := 1
	ans := 0
	
	// 0 1 1 2 3 5 8 13 21 34

	for i := 0; i < a; i++ {
		ans = n1 + n2
	
		n1 = n2
		n2 = ans
	
		
	
	}
	fmt.Println("fibonacci num =>",ans)
	
}