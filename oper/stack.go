package main

import (
	"log"
)

// Stack ...
type Stack struct {
	Top  *Node // "Top"
	Size int   // "sd"
}

// Node ...
type Node struct {
	Value interface{} //
	Next  *Node       //
}

func (st *Stack) show() {
	for st.Size > 0 {
		log.Println(st.pop().Value)
	}
}

func inits() Stack {
	log.Println("inits")
	st := Stack{Top: &Node{}, Size: 0}
	return st
}

func (st *Stack) pop() Node {
	st.Size--
	res := st.Top
	st.Top = st.Top.Next
	return *res
}

func (st *Stack) push(v interface{}) {
	n := Node{v, st.Top}
	st.Top = &n
	st.Size++
}

// Parse ...
func Parse(v int) {
	s := inits()
	exe(v, &s)
	s.show()
}

func exe(v int, s *Stack) Stack {
	if v > 0 {
		t := v % 8
		v = v / 8
		s.push(t)
		exe(v, s)
	}
	return *s
}

/* func main() {
	fmt.Println("test Stack")
	s := inits()
	log.Println(s.Size)
	s.push(1)
	s.push(2)
	s.push(3)
	s.show()
	fmt.Println("test int10 -> int8")
	v := 23
	Parse(v)
}
*/
