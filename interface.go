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

//	rect -> joins rectangle struct -> rect is used as an argument for calcArea -> rect is accepted as a valid argument
//	since its just two float64 value's -> rect is passed to the nested area function which belongs to the Shape
//	interface -> rect gets sent to the version of the function for its respective struct -> the function multiplies
//	the two value's belonging to rect -> this sucsesfully returns a float64 value meeting the criteria of the shape
// interface.
