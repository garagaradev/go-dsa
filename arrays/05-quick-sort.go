package main
import "fmt"

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    
    for j := low; j < high; j++{
    if arr[j] <= pivot {
      i++

      arr[i],arr[j] = arr[j], arr[i]
    }
  }

  arr[i+1], arr[high] = arr[high],arr[i+1]

  return i + 1
}

func quicksort(arr []int, low, high int){
  if low < high {
    pivotIndex := partition(arr, low, high)

    quicksort(arr, low, pivotIndex-1)
    quicksort(arr, pivotIndex+1,high)
  }
}

func main(){
  data := []int{64,34,25,12,22,11,90,5}
  fmt.Println("Before:",data)
  quicksort(data,0, len(data)-1)
  fmt.Println("After:",data)
}
