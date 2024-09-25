package main
import "fmt"


type Stack struct {
  stack []string
}

// push a method to add an element to the stack  (pointer receiver)
func (s *Stack) Push(element string) {
  s.stack = append(s.stack, element)
}

// pop method to remove and return the top element of the stack 
func (s *Stack) Pop() string {
  if s.IsEmpty(){
    return "Stack is empty"
  }

  //get the last element 
  element := s.stack[len(s.stack)-1]
  // remove the last element from the stack 
  s.stack = s.stack[:len(s.stack)-1]
  return element
}


// peek method to return the top element without removing it
func (s *Stack) Peek() string {
  if s.IsEmpty(){
    return "Stack is empty"
  }

  return s.stack[len(s.stack)-1]
}

// IsEmpty method to check if the stack is empty
func (s *Stack) IsEmpty() bool {
  return len(s.stack) == 0
}

// size method to return the number of elements in the stack 
func (s *Stack) Size() int {
  return len(s.stack)
}

func main(){
  // Create a stack 
  myStack := &Stack{}

  myStack.Push("A")
  myStack.Push("B")
  myStack.Push("C")
  fmt.Println("Stack:", myStack.stack)
  fmt.Println("Pop:", myStack.Pop())
  fmt.Println("Peek:", myStack.Peek())
  fmt.Println("IsEmpty:", myStack.IsEmpty())
  fmt.Println("Size:", myStack.Size())
}
