package main

import "fmt"

func main() {
	s := Stack{}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println(s.Pop())     // Should print: 30, true
	fmt.Println(s.Peek())    // Should print: 20, true
	fmt.Println(s.IsEmpty()) // Should print: false
}
