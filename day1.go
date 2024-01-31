
package main
import "fmt"

func main(){
	// // lets first print Hello World
	// fmt.Println("Hello World")

	// // variable declaration
	// number := 1
	// fmt.Println("vARIABLE 3", number)

	// // if statements in go lang
	// item1 := 3
	// item2 := 4
	// if item1 > item2 {
	// 	fmt.Println(false)
	// }
	// fmt.Println(true)

	// for loop in Go
	// for i := 5; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// Go data structures
	// Array is a group of similar types of data

	numbers := [6]int{7, 9, 1, 2, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])

	}
		


}



