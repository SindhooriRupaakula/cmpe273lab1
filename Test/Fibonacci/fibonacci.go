package fibonacci

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
