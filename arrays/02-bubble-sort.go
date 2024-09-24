package main
import "fmt"


func main(){

  var data [8]int
  data = [8]int{64, 34, 25, 12, 22, 11, 90, 5}
  fmt.Println("before:",data)

  n := len(data)
  for i := range n-1 {
    for j := range n-i-1 {
      if data[j] > data[j+1]{
        data[j],data[j+1]= data[j+1], data[j]
      }
    }
  } 
  fmt.Println("after:",data)

}
