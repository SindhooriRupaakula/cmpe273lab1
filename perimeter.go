package main
import ("fmt" ; "math")

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


func main(){
  var r Rectangle
  var c Circle

  fmt.Println("Enter length and width of rectangle: ")
  fmt.Scan(&r.l, &r.w)
  fmt.Printf("Perimeter of rectangle is: %.2f \n", r.perimeter())

  fmt.Println("Enter radius of circle: ")
  fmt.Scan(&c.r)
  fmt.Printf("Perimeter of circle is: %.2f \n", c.perimeter())
  fmt.Println(c.perimeter())
}
