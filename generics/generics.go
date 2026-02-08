package main
import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func add[t int | float64](x t, y t) t { //	Instead of having a set variable type we are declaring t(type) to accept
																				//	either a int or float64 value.
	return x + y
}
//	This is still pretty limitied in scope for all the different types of integer values. a better way of doing this
//	is:

type numTypes interface { 								//  You can also use type name int or type name int | int8 | etc. which is
	int | int8 | int16 | float32 | float64	//  can be used in a func likeso: x := name(value). to pass that into a func
}																					//  make sure the variable type includes a tilda infront of likeso:
//	And then;															//  func add[t ~int](x t, y t) t {}

func addInterface[t numTypes](x t, y t) t {
	return x + y
}
//	In an actual project you can import the constraints library which is the previous interface solution with all of
//	boiler plate taken care of.
//
//	It is still probably a good idea to know what the library is doing for you though which is why the previous
//	example was shown.

func addConstraints[t constraints.Ordered](x t, y t) t {
	return x + y
}

func main() {
	x, y := 1.5, 5.5

	fmt.Println(addConstraints(x, y))
}
