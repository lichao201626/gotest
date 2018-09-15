package main

import (
	"log"
	"testing"
)

func Test_str(t *testing.T) {
	str := "1 2+3*4*2-5/1+3-2"
	res := eval(str)
	log.Println(res)
	/* 	for k, v := range res {
		log.Println("range res", k, v)
		t.Log("something", k, v)
	} */

}
