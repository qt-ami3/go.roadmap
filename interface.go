//Flowchart
//	rect -> rectangle struct (x, y) -> 

package main
import (
	"math"
	"fmt"
)

//'A': Public		'a': private
type Shape interface {
	area() float64
}

type rectangle struct {
	width, height		float64
}

type circle struct {
	radius		float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calcArea(s Shape) float64 {
	return s.area()
}

func main() {
	rect := rectangle{width: 5, height: 4}
	circle := circle{radius: 2}

	fmt.Println(
		"rectangle area: ",
		calcArea(rect),
	)
	
	fmt.Println(
		"circle area: ",
		calcArea(circle),
	)
}
