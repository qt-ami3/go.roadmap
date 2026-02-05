package main
import "fmt"

func sum(x, y int)(result int) {
	result = x + y //dont need ":=" because result int is a variable declaration.
	return result
}

func main() {
	fmt.Println(sum(1,1))
}
