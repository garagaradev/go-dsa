package main
import "fmt"

func main(){
  //create array
  var a [10]int
  a[0] = 10
  a[1] = 5
  fmt.Println(a)

  for i:=2;i<len(a);i++{
    a[i] += i*30
  }
  fmt.Println(a)
  a[5] = 33
  fmt.Println(a)
}

