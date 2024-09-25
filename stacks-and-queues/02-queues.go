package main
import "fmt"

type Queue struct {
  queue []string
}

// Enqueue method (to add an element to the queue)
func (q *Queue) Enqueue(element string){
  q.queue = append(q.queue, element)
}

// Dequeue method (to remove and return the front element of the queue)
func (q *Queue) Dequeue() string{
  if q.IsEmpty(){
    return "Queue is empty"
  }

  //get the front element
  element := q.queue[0]
  //remove the front element
  q.queue = q.queue[1:]
  return element
}

// peek method (to return the front element without removing it)
func (q *Queue) Peek() string{
  if q.IsEmpty(){
    return "Queue is empty"
  }
  return q.queue[0]
}

func (q *Queue) IsEmpty() bool {
  return len(q.queue) == 0
}

func (q *Queue) Size() int {
  return len(q.queue)
}

func main(){
  // create a queue
  myQueue := &Queue{}

  myQueue.Enqueue("A")
  myQueue.Enqueue("B")
  myQueue.Enqueue("C")
  
  fmt.Println("Queue:", myQueue.queue)
  fmt.Println("Dequeue:", myQueue.Dequeue())
  fmt.Println("Peek:",myQueue.Peek())
  fmt.Println("IsEmpty:", myQueue.IsEmpty())
  fmt.Println("Size:", myQueue.Size())
}
