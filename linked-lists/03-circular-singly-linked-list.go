package main
import "fmt"

type Node struct {
  data int
  next *Node
}

func main(){
  //create nodes
  node1 := &Node{data:3}
  node2 := &Node{data:5}
  node3 := &Node{data:13}
  node4 := &Node{data:2}

  // link nodes in a circular style 
  node1.next = node2
  node2.next = node3
  node3.next = node4
  node4.next = node1

  // traverse 
  currentNode := node1
  startNode := node1

  fmt.Printf("%d -> ", currentNode.data)
  currentNode = currentNode.next


  for currentNode != startNode {
    fmt.Printf("%d -> ", currentNode.data)
    currentNode = currentNode.next
  }

  fmt.Println("...") // indicator that the circular singly linked list happened
}
