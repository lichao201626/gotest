package regexmatch

import (
	"github.com/lichao201626/gotest/server/letcode/regularexpressionmatching"
	"testing"
)

func Test_regxmatch(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		res := regexmatch.IsMatch(v.ind1, v.ind2)
		if res != v.outd {
			t.Error("Case failed", v.ind1, v.ind2, v.outd)
		} else {
			t.Log("Case passed", v.ind1, v.ind2, v.outd)
		}
	}
}

type data struct {
	ind1 string
	ind2 string
	outd bool
}

func initD() []data {
	res := []data{{ind1: "ab", ind2: ".*c", outd: false},
		{ind1: "aaaaaaaa", ind2: "a.aaa.a.a", outd: false},
		{ind1: "aaaaaaaa", ind2: "a.", outd: false},
		{ind1: "", ind2: ".*", outd: true},
		{ind1: "aaaaaaaa", ind2: "**a*", outd: true},
		{ind1: "aab", ind2: "c*a*b", outd: true},
		{ind1: "mississippi", ind2: "mis*is*p*.", outd: false},
		{ind1: "mississippi", ind2: "mis*is*ip*.", outd: true},
	}
	return res
}
