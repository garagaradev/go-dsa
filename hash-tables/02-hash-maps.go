package main
import (
  "fmt"
  "strconv"
)

// SimpleHashMap represents a simple hash map
type SimpleHashMap struct {
  size int
  buckets [][]KeyValuePair // array of slices to handle collisions
}

type KeyValuePair struct {
  key string
  value string
}

// New SimpleHashMap inits a new hash map with the given size
func NewSimpleHashMap(size int) *SimpleHashMap {
  buckets := make([][]KeyValuePair, size)
  for i := range buckets {
    buckets[i] = []KeyValuePair{}
  }

  return &SimpleHashMap{
    size: size,
    buckets: buckets,
  }
}

// HashFunction computes a simple hash based on the sum of numeric characters modulo 10 (in this case)
func (h *SimpleHashMap) HashFunction(key string) int {
  sum := 0
  for _, char := range key {
    if n, err := strconv.Atoi(string(char)); err == nil {
      sum += n
    }
  }

  return sum % 10
}

// Put inserts or updates a key-value pair in the hash map
func (h *SimpleHashMap) Put(key string, value string) {
  index := h.HashFunction(key)
  bucket := h.buckets[index]
  for i, pair := range bucket {
    if pair.key == key {
      // update existing key
      h.buckets[index][i].value = value
      return
    }
  }

  // add new key-value pair
  h.buckets[index] = append(bucket, KeyValuePair{key, value})
}

// Get method retrieves  a value by key 
func (h *SimpleHashMap) Get(key string) string {
  index := h.HashFunction(key)
  bucket := h.buckets[index]
  for _, pair := range bucket {
    if pair.key == key {
      return pair.value
    }
  }
  return "" // key not found
}

// 'Remove method' deletes a key-value pair by key
func (h *SimpleHashMap) Remove(key string){
  index := h.HashFunction(key)
  bucket := h.buckets[index]
  for i, pair := range bucket {
    if pair.key == key {
      // remove the key-value pair
      h.buckets[index] = append(bucket[:i],bucket[i+1:]...)
      return
    }
  }
}

// 'PrintMap method' prints all key-value pairs in the hashmap
func (h *SimpleHashMap) PrintMap(){
  fmt.Println("Hash Map Contents:")
  for i, bucket := range h.buckets {
    fmt.Printf("Bucket %d: %v\n",i, bucket)
  }
}


func main(){
  // create a new hash map
  hashMap := NewSimpleHashMap(10)

  // adding some entries
  hashMap.Put("123-4567","Charlotte")
  hashMap.Put("123-4568","Thomas")
  hashMap.Put("123-4569","Jens")
  hashMap.Put("123-4570","Peter")
  hashMap.Put("123-4571","Lisa")
  hashMap.Put("123-4572","Adele")
  hashMap.Put("123-4573","Michaela")
  hashMap.Put("123-6574","Bob")

  hashMap.PrintMap()

  //retrieval
  fmt.Println("\nName associated with '123-4570':",hashMap.Get("123-4570"))
  //updating the name
  fmt.Println("Updating the name for '123-4570' to James")
  hashMap.Put("123-4570","James")
  //checking Peter's availability
  fmt.Println("Name associated with '123-4570':", hashMap.Get("123-4570"))

}
