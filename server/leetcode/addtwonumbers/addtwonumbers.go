package addtwonumbers

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

/*AddTwoNumbers  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807. */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	var d1, d2 int
	temp := 0
	tempN := res
	for l1 != nil || l2 != nil {
		if l1 != nil {
			d1 = l1.Val
			l1 = l1.Next
		} else {
			d1 = 0
		}
		if l2 != nil {
			d2 = l2.Val
			l2 = l2.Next
		} else {
			d2 = 0
		}
		v := d1 + d2
		val := v % 10
		temp = v / 10
		if temp > 0 {
			if l1 != nil {
				l1.Val = l1.Val + temp
			} else if l2 != nil {
				l2.Val = l2.Val + temp
			} else {
				l2 = &ListNode{
					Val:  temp,
					Next: nil,
				}
			}
		}
		n := ListNode{
			Val:  val,
			Next: nil,
		}
		tempN.Next = &n
		tempN = tempN.Next
	}
	return res.Next
}

//Ap ...
func (ln *ListNode) Ap(v int) {
	i := 1
	for ln.Next != nil {
		i++
	}
	node := ListNode{
		Val:  v,
		Next: &ListNode{},
	}
	ln.Next = &node
}

func arrayToList(arr []int) ListNode {
	list := ListNode{}
	for i := len(arr) - 1; i >= 0; i-- {
		// fmt.Print(arr[i])
		list.Ap(arr[i])
	}
	return list
}
