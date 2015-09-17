package perimeter
import "testing"

func Test_perimeter(t *testing.T) {
 r := Rectangle{23, 22.5}
 c := Circle{10}

 if r.perimeter() != 45.50 {
  t.Error("For perimeter of rectangle with length ", r.l, "and width", r.w, "expected", 45.50,"got", r.perimeter)
 }

 if c.perimeter() != 314.1592653589793 {
  t.Error("For perimeter of circle with radius ", c.r, "expected", 45.50,"got", r.perimeter)
 }

}
