package main
import "fmt"

//function to calculate fibbonacci numbers
func fib(n int) int{
  if n == 0 {
    return n
  } else if n == 1 {
    return 1
  } else {
    return fib(n - 1) + fib(n - 2)
  }
}

func main(){
        var num int

        fmt.Println("Please input the number of fibonacci numbers to be printed: ")
        fmt.Scan(&num)
        if num < 0 {
          fmt.Println("\n\n!!!You have entered a negative number. Please try again later!!!")
        } else {
          fmt.Println("\n\nPrinting", num, "fibonacci numbers: ")
          for i := 0; i<num ; i++ {
            fmt.Println(fib(i), " ")
          }
        }
}
