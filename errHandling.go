package main

import (
	"fmt"
	"errors"
)

func doSomething() error {
	return errors.New("something broke lol")
}

func divide(x, y int) (int, error) {	//	The function is expecting two values an int & an error value.
	if y == 0 {
		return 0, fmt.Errorf("cannot divide '%d' by zero", x)	//  0, error
	}
	return x / y, nil	//  x / y, no error
}

func main() {
	x := 10
	y := 0
	z := 2
	
	fmt.Println("10 / 0:")
		fmt.Println(divide(x, y))
	fmt.Println()

	fmt.Println("10 / 2:")
		fmt.Println(divide(x, z))
}
