package fact

import "fmt"

func Fact(a int) {

	// factorial

	var ans int = 1
	
	for i := 0; i < a; i++ {
		ans *= a - i
	
	}
	
	fmt.Println("factorial => ", ans)
	
	
}