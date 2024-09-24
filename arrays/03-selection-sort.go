package main
import "fmt"

func main(){
  data := []int{64,34,25,5,22,11,90,12}
  fmt.Println("Before:",data)

  n := len(data)

  for i := range n {
    minIndex := i
    for j:=i+1; j < n; j++ {
      if data[j] < data[minIndex]{
        minIndex = j
      }
    }
    minValue := data[minIndex]

    data = append(data[:minIndex], data[minIndex+1:]...)
    data = append(data[:i],append([]int{minValue}, data[i:]...)...)
  }

  fmt.Println("After:",data)
}

