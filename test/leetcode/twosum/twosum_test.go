package twosum

import (
	"testing"

	"github.com/lichao201626/gotest/server/leetcode/twosum"
)

func Test_twosum(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		res := twosum.TwoSum(v.ind1, v.ind2)
		if (res[0] == v.outd[0] && res[1] == v.outd[1]) ||
			(res[0] == v.outd[1] && res[1] == v.outd[0]) {
			t.Log("Case passed", v.ind1, v.ind2, v.outd)
		} else {
			t.Error("Case failed", v.ind1, v.ind2, v.outd)
		}

	}
}

type data struct {
	ind1 []int
	ind2 int
	outd []int
}

func initD() []data {
	res := []data{{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 3}, 6, []int{0, 2}},
	}
	return res
}
