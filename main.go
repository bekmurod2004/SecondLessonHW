package main


import ("fmt"
"lessonGo/seconLesson/SecondLessonHW/fact"
"lessonGo/seconLesson/SecondLessonHW/fib"
"lessonGo/seconLesson/SecondLessonHW/fiz"

)

func main() {
	var number int 
	fmt.Scanln(&number)
	fmt.Println()

	fiz.Fizbuzer(number)
	fmt.Println()

	fib.Fib(number)
	fmt.Println()
	
	fact.Fact(number)
}