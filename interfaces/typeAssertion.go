//Type assertions add overhead and increase cpu utilization so it should be avoided for a more performent solution in 
//looping or reoccuring logic.
//
//What this is really helpful for is converting the value of a variable to another compatible data type.

package main
import "fmt"

func main() {
	//	These are represent interface and are not the same as type inference (:=)
	var x interface{} = "Hello" //	or var y any = "Hello"
	
//bin := x.(string) //	Statically assign current value of interface to the variable but will panic if var type is not
										//	as expected.
	
	if bin, ok := x.(string); ok { //	Wont panic if variable type is not as expected.
		fmt.Println(bin)
	}
}
