package myfs

import (
	"github.com/lichao201626/gotest/server/myfs"
	"testing"
)

func Test_one(t *testing.T) {
	/* 	for i := 0; i < 5; i++ {
	   		t.Fatalf("as")
	   	}
	   	t.Fail()
	   	t.Failed()
		   t.Error("fail") */
	myfs.WriteF("a/b/c/name.sx", "asdfal ,sdjf 中国")
	t.Log("asdf")
}
