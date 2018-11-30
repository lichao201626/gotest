package addtwonumbers

import (
	"fmt"
	"github.com/lichao201626/gotest/server/leetcode/addtwonumbers"
	"testing"
)

func Test_addtwonumbers(t *testing.T) {
	mock := initD()
	for k, v := range mock {
		t.Log("asdf", k, v)
		in1 := arrayToList(v.ind1)
		in2 := arrayToList(v.ind2)
		res := addtwonumbers.AddTwoNumbers(in1, in2)
		if isPass(*res, v.outd) {
			t.Log("Case passed", v.ind1, v.ind2, v.outd)
		} else {
			t.Error("Case failed", v.ind1, v.ind2, v.outd)
		}
	}
}

type data struct {
	ind1 []int
	ind2 []int
	outd []int
}

func initD() []data {
	res := []data{{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{5}, []int{5}, []int{0, 1}},
	}
	return res
}

func arrayToList(arr []int) *addtwonumbers.ListNode {
	head := new(addtwonumbers.ListNode)
	tp := head
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		tp.Next = &addtwonumbers.ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tp = tp.Next
	}
	return head.Next
}

func isPass(ln addtwonumbers.ListNode, arr []int) bool {
	isPass := true
	for i := 0; i < len(arr); i++ {
		if arr[i] != ln.Val {
			isPass = false
			break
		}
		if i != len(arr)-1 {
			ln = *ln.Next
		}

	}
	return isPass
}
