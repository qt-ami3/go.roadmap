package main
import "fmt"

func describe(i interface{}) { //You can always use func describe(i... interface{}) {} to pass multiple arguments.
	switch v := i.(type){ //v for variable
		case int:
			fmt.Println("this is an int: ", v)
		case string:
			fmt.Println("this is an string: ", v)	
		default:	
			fmt.Println("this is an a mysterious other thing: ", v)
	}
}

func main() {
	describe(5)
	describe("Hello World!")
	describe(false)
}

//	Empty interfaces can be helpful when you dont know what kind of data type you are working with.

