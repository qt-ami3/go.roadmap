//its a function that can be run without being defined, crazy stuff.

package main
import "fmt"

func main() {
	func() {
		fmt.Println("Hello World!")
	}()

	addedValue := func(x, y int) int {
		return x + y
	}

	fmt.Println(addedValue(5, 7))
}
