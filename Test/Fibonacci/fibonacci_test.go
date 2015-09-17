package fibonacci
import "testing"

func Test_fib (t *testing.T) {
 v0 := fib(0)
 v1 := fib(1)
 v5 := fib(5)
 v10 := fib(10)

 if v0 != 0 {
   t.Error("For fib(0)", "expected", 0,"got", v0)
 }

 if v1 != 1 {
   t.Error("For fib(1)", "expected", 1,"got", v1)
 }

 if v5 != 5 {
   t.Error("For fib(5)", "expected", 5,"got", v5)
 }

 if v10 != 55 {
   t.Error("For fib(10)", "expected", 55,"got", v10)
 }

}
