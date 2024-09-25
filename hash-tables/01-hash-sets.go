package main
import "fmt"

type SimpleHashSet struct {
  size int
  buckets [][]string // array of slices to handle collisions
}
//a function to init a new hash set with the given size
func NewSimpleHashSet(size int) *SimpleHashSet {
  buckets := make([][]string, size)
  for i := range buckets {
    buckets[i] = []string{}
  }

  return &SimpleHashSet{
    size: size,
    buckets: buckets,
  }
}


// hash fx computes a simple hash based on the sum of character codes module the number of buckets
func (h *SimpleHashSet) HashFunction(value string) int {
  sum := 0
  for _, char := range value {
    sum += int(char)
  }

  return sum % h.size
}

// Add inserts a value into the hash set if it's not already present
func (h *SimpleHashSet) Add(value string){
  index := h.HashFunction(value)
  bucket := h.buckets[index]
  for _, v := range bucket {
    if v == value {
      return
    }
  }

  h.buckets[index] = append(h.buckets[index],value)
}

// contains checks if a value exists in the hash set
func (h *SimpleHashSet) Contains(value string) bool {
  index := h.HashFunction(value)
  bucket := h.buckets[index]
  for _, v := range bucket {
    if v == value {
      return true
    }
  }

  return false
}

// remove method deletes a value from the hash set if it exists
func (h *SimpleHashSet) Remove(value string){
  index := h.HashFunction(value)
  bucket := h.buckets[index]
  for i, v := range bucket {
    if v == value {
      h.buckets[index] = append(bucket[:i], bucket[i+1:]...)
      return
    }
  }
}


// PrintSet method prints all elements in the hash set
func (h *SimpleHashSet) PrintSet(){
  fmt.Println("Hash Set Contents:")
  for i, bucket := range h.buckets {
    fmt.Printf("Bucket %d: %v\n", i, bucket)
  }
}

func main(){
  //create a new hash set
  hashSet := NewSimpleHashSet(10)
  // Add elements to the hash set
  hashSet.Add("Charlotte")
  hashSet.Add("Thomas")
  hashSet.Add("Jens")
  hashSet.Add("Peter")
  hashSet.Add("Lisa")
  hashSet.Add("Adele")
  hashSet.Add("Michaela")
  hashSet.Add("Bob")

  hashSet.PrintSet()

  fmt.Println("\n'Peter' is in the set?", hashSet.Contains("Peter"))
  fmt.Println("Removing Peter...")
  hashSet.Remove("Peter")
  fmt.Println("'Peter' is in the set?",hashSet.Contains("Peter"))

  fmt.Println("'Adele' has the hash code:",hashSet.HashFunction("Adele"))
}

