package myatoi

import (
	"github.com/lichao201626/gotest/server/leetcode/ispalindrome"
	"testing"
)

func Test_atoi(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		res := ispalindrome.Ispalindrome(v.ind)
		if res != v.outd {
			t.Error("Case failed", v.ind, v.outd)
		} else {
			t.Log("Case passed", v.ind, v.outd)
		}
	}
}

type data struct {
	ind  int
	outd bool
}

func initD() []data {
	res := []data{{ind: 121, outd: true},
		{ind: -121, outd: false},
		{ind: 11, outd: true},
		{ind: 0, outd: true},
	}
	return res
}
