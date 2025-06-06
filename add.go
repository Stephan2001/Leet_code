package add

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := &ListNode{Val: 0}
	return Walk(l1, l2, a, 0)
}

func Walk(n1, n2, nres *ListNode, remainder int) *ListNode {
	v1, v2 := 0, 0
	if n1 != nil {
		v1 = n1.Val
	}
	if n2 != nil {
		v2 = n2.Val
	}

	// Sum + carry
	sum := v1 + v2 + remainder
	newRemainder := 0
	if sum > 9 {
		sum -= 10
		newRemainder = 1
	}

	nres.Val = sum

	more1 := (n1 != nil && n1.Next != nil)
	more2 := (n2 != nil && n2.Next != nil)
	if more1 || more2 || newRemainder > 0 {
		nres.Next = &ListNode{}

		var next1, next2 *ListNode
		if n1 != nil {
			next1 = n1.Next
		}
		if n2 != nil {
			next2 = n2.Next
		}

		Walk(next1, next2, nres.Next, newRemainder)
	}

	return nres
}
