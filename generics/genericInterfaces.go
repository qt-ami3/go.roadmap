package main
import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type customData interface {
	constraints.Ordered | []byte | []rune
}

type Usr[t customData] struct {
	id		int
	name	string
	data	t
}

func main() {
	u := Usr[int]{
		id:		0,
		name:	"jim",
		data:	3,
	}

	fmt.Println(u)
}
