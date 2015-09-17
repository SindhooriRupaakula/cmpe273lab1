package perimeter
import "math"

type Shape interface {
 perimeter() float64
}

//rectangle with coordinates
type Rectangle struct {
  l, w float64
}

//circle with radius
type Circle struct {
  r float64
}

func (rect *Rectangle) perimeter() float64 {
  return rect.l + rect.w
}

func (circ *Circle) perimeter() float64 {
  return math.Pi * circ.r * circ.r
}
