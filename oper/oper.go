package main

import (
	"fmt"
	"log"
)

// const of oper
const (
	Plus     = "+"
	Minus    = "-"
	Multiple = "*"
	Divide   = "/"
)

func main() {
	fmt.Println("test oper")
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
