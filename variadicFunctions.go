//vah-re-ah-dic
//a function that takes an undefined amount of arguments as input:
//fmt.Println("string", stringVariable)

package main
import "fmt"

func undefIntMultiply(nums... int) int {
	var sum int = 1
	for _, num := range nums {
		sum  *= num
	}
	return sum
}

func main() {
	arr := []int{2,3,7}
	fmt.Println(undefIntMultiply(arr...))
	fmt.Println(undefIntMultiply(1, 5, 5))
}
