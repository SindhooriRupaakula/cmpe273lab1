package main
import ( "fmt"
         "time" )

//Funtion for sleep
func mySleep(t int){
 fmt.Println("\n Going to Sleep...")

 <- time.After(time.Second * time.Duration(t))
 fmt.Println("\n System Ready!")
}

func main(){
  var t int
  fmt.Println("\n Please enter sleep time in seconds: ")
  fmt.Scan(&t)
  mySleep(t)
}
