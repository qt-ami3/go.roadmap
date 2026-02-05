package main
import "fmt"

func (c *Child) requestMoney() int{
	fm := c.f.giveMoney()
	return fm + c.savings
}

func (f *Father) giveMoney() int{
	return 100
}

type Child struct {
	savings int
	f Father
}

type Father struct {	
//	empty struct
}

func main() {
	f := Father{}
	c := Child{savings: 50, f:f}	//	c.savings = 50

	fmt.Println(c.requestMoney())	//	giveMoney of Father of Child returns an int value of 100 and adds it to fm
																//	which is returned as itself plus savings of Child as an int value which itself is
																//	printed.
}
