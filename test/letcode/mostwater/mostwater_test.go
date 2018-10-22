package mostwater

import (
	"github.com/lichao201626/gotest/server/letcode/mostwater"
	"testing"
)

func Test_mostwater(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		res := mostwater.MaxArea(v.ind1)
		if res != v.outd {
			t.Error("Case failed", v.ind1, v.outd)
		} else {
			t.Log("Case passed", v.ind1, v.outd)
		}
	}
}

type data struct {
	ind1 []int
	outd int
}

func initD() []data {
	res := []data{{ind1: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, outd: 49}}
	return res
}
