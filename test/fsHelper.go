package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func writeF(ff string, content interface{}) bool {
	// var ff = "a/b.abc"
	w := strings.Split(ff, "/")
	for k, v := range w {
		log.Println(k, v)
	}
	var rdNum = rand.Intn(100)
	log.Println("asdfasdf", rdNum)
	// bytes := make([]byte, 10)
	x := os.ModePerm
	fn := "fileName"
	fn += strconv.Itoa(rdNum)
	fn += ".abc"
	/* 	dir := "a"
	   	sp := filepath.Separator
	   	y := strconv.QuoteRune(sp)
	   	log.Println("asdfasdf", y)
	   	var fullName = dir + y + fn
	   	e := ioutil.WriteFile(fullName, bytes, x)
	   	if e != nil {
	   		log.Println(e)
		   } */

	bts := []byte{}
	switch content.(type) {
	/* 	case int:
	log.Println(content, "is int") */
	case string:
		bts = []byte(content.(string))
		log.Println(content, "is string")
		/* 	case float64:
		   		log.Println(content, "is float64")
		   	case bool:
		   		log.Println(content, " is bool") */
	default:
		log.Println("未知的类型")

	}
	// content := []byte("temporary file's content")
	var dir = "a"
	err := os.Mkdir(dir, x)
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, fn)
	log.Println(tmpfn)
	if err := ioutil.WriteFile(tmpfn, bts, x); err != nil {
		log.Fatal(err)
	}
	return true
}

func main() {
	writeF("a/test.ls", "hello world")
}
