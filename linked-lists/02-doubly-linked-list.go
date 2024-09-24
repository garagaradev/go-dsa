package main
import "fmt"

type Node struct {
  data int
  next *Node
  prev *Node
}

func main(){
  node1:= &Node{data:3}
  node2:= &Node{data:5}
  node3:= &Node{data:13}
  node4:= &Node{data:2}

  node1.next = node2

  node2.prev = node1
  node2.next = node3

  node3.prev = node2
  node3.next = node4

  node4.prev = node3

  fmt.Println("\n Transversing forward:")
  currentNode := node1

  for currentNode != nil {
    fmt.Printf("%d ->", currentNode.data)
    currentNode = currentNode.next
  }
  fmt.Println("null")

  fmt.Println("\n Transversing backward:")
  currentNode = node4
  for currentNode != nil {
    fmt.Printf("%d ->", currentNode.data)
    currentNode = currentNode.prev
  }
  fmt.Println("null")
}
