package myatoi

import (
	"github.com/lichao201626/gotest/server/leetcode/myatoi"
	"testing"
)

func Test_atoi(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		res := myatoi.Myatoi(v.ind)
		if res != v.outd {
			t.Error("Case failed", v.ind, v.outd)
		} else {
			t.Log("Case passed", v.ind, v.outd)
		}
	}
}

type data struct {
	ind  string
	outd int
}

func initD() []data {
	res := []data{{ind: "42", outd: 42},
		{ind: " -42", outd: -42},
		{ind: "4193 with words", outd: 4193},
		{ind: "words and 987", outd: 0},
		{ind: "-91283472332", outd: -2147483648},
		{ind: "9223372036854775808", outd: 2147483647},
		{ind: "", outd: 0},
	}
	return res
}
