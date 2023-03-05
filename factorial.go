package main

import "fmt"

func fact(a int) {

	// factorial

	var ans int = 0
	
	for i := 0; i < a; i++ {
		ans += a - i
	
	}
	
	fmt.Println("factorial => ", ans)
	
	
}