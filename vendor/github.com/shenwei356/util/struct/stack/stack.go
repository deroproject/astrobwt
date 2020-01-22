package stack

// Stack struct
type Stack []interface{}

// Empty tell if it is empty
func (s Stack) Empty() bool {
	return len(s) == 0
}

// Peek return the last element
func (s Stack) Peek() interface{} {
	return s[len(s)-1]
}

// Put puts element to stack
func (s *Stack) Put(i interface{}) {
	(*s) = append((*s), i)
}

// Pop pops element from the stack
func (s *Stack) Pop() interface{} {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

/*
func main() {
	var s stack

	for i := 0; i < 3; i++ {
		s.Put(i)
		fmt.Printf("len=%d\n", len(s))
		fmt.Printf("peek=%d\n", s.Peek())
	}

	for !s.Empty() {
		i := s.Pop()
		fmt.Printf("len=%d\n", len(s))
		fmt.Printf("pop=%d\n", i)
	}
}
*/
