package main
import "fmt"

func main(){
  data := []int{64, 34, 25, 12, 22, 11, 90, 5}
  fmt.Println("Before:",data)

  n := len(data)

  for i:=1;i<n; i++{
    insertIndex := i
    currentValue := data[i]
    data = append(data[:i], data[i+1:]...)

    for j:=i-1; j >= 0; j-- {
      if data[j] > currentValue {
        insertIndex = j
      }else{
        break
      }
    }

    data = append(data[:insertIndex], append([]int{currentValue}, data[insertIndex:]...)...)
  }

  fmt.Println("After:", data)
}
