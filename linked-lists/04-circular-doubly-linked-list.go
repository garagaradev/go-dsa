package main
import "fmt"


type Node struct {
  data int
  next *Node
  prev *Node
}

func main(){
  //create the nodes, guys
  node1 := &Node{data: 3}
  node2 := &Node{data: 5}
  node3 := &Node{data: 13}
  node4 := &Node{data: 2}

  // link the nodes forward and backward. Oh yeah
  node1.next = node2
  node1.prev = node4

  node2.next = node3
  node2.prev = node1

  node3.next = node4
  node3.prev = node2

  node4.next = node1
  node4.prev = node3

  fmt.Println("\nTraversing forward:")
  currentNode := node1
  startNode := node1
  fmt.Printf("%d -> ", currentNode.data)
  currentNode = currentNode.next

  for currentNode != startNode {
    fmt.Printf("%d -> ", currentNode.data)
    currentNode = currentNode.next
  }

  fmt.Println("...")

  fmt.Println("\nTraversing backward:")
  currentNode = node4
  startNode = node4
  fmt.Printf("%d -> ",currentNode.data)
  currentNode = currentNode.prev

  for currentNode != startNode {
    fmt.Printf("%d -> ", currentNode.data)
    currentNode = currentNode.prev
  }

  fmt.Println("...")

}
