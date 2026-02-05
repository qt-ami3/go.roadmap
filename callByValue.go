//Call by value:
//	Pass the value of the parameter/argument making any changes not effect the original value.
//
//Call by reference:
//	Pass the memory address of the parameter.
//	Since you are using the memory address as an argument, any changes inside of the function will be done to the
//	variable used as the argument.

package main
import "fmt"


//Call by reference function:
func initFloor(floor *int) {
	//psuedo code
}

//Call by value function:
func initFloor(floor int) {
	//psuedo code
}

func main() {
	//Call by reference:
	initFloor(&floor)

	//Call by value:
	initFloor(floor)
}
