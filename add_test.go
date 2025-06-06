package add

import "testing"

func TestWalkProgress(t *testing.T) {

	node1 := &ListNode{Val: 0}
	node2 := &ListNode{Val: 9}
	node3 := &ListNode{Val: 9}
	node4 := &ListNode{Val: 9}
	node5 := &ListNode{Val: 9}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node6 := &ListNode{Val: 0}
	node7 := &ListNode{Val: 9}
	node8 := &ListNode{Val: 9}
	node6.Next = node7
	node7.Next = node8

	dummy := &ListNode{Val: 0}

	dummy = Walk(node1, node6, dummy, 0)

	p1 := node1
	p2 := node4
	pr := dummy

	for p1 != nil || p2 != nil || pr != nil {
		var v1, v2, rv int
		if p1 != nil {
			v1 = p1.Val
		}
		if p2 != nil {
			v2 = p2.Val
		}
		if pr != nil {
			rv = pr.Val
		}

		t.Logf("input1: %d, input2: %d → result : %d", v1, v2, rv)

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
		if pr != nil {
			pr = pr.Next
		}
	}
}

func Walker() *ListNode {
	a := &ListNode{Val: 0}
	b := &ListNode{Val: 7}
	c := &ListNode{Val: 0}
	d := &ListNode{Val: 8}

	a.Next = b
	b.Next = c
	c.Next = d

	return a
}
