//	goal:
//
//	input: [1, 2, 3]
//	function output = input * 2

package main
import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func nonGeneric(values []int, mapFunc func(int) int) []int {
	var bin []int
	for _, v := range values {
		binTwo := mapFunc(v)
		bin = append(bin, binTwo)
	}
	return bin
}

func generic[t constraints.Ordered](values []t, mapFunc func(t) t) []t {
	var bin []t
	for _, v := range values {
		binTwo := mapFunc(v)
		bin = append(bin, binTwo)
	}
	return bin
}

func main() {
	result := generic([]float64{1.5, 2.5, 3.5}, func(n float64) float64 {
		return n * 2
	})

	fmt.Println(result)
}
