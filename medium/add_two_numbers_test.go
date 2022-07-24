package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 2. Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Constraints:
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

func Benchmark_Test_Add_Two_Numbers_1(b *testing.B) {
	listNodeL1 := ListNode{3, nil}
	listNodeL2 := ListNode{4, &listNodeL1}
	headNodeL := ListNode{2, &listNodeL2}

	listNodeR1 := ListNode{4, nil}
	listNodeR2 := ListNode{6, &listNodeR1}
	headNodeR := ListNode{5, &listNodeR2}

	resultNode := AddTwoNumbers(&headNodeL, &headNodeR)

	expectNode1 := ListNode{8, nil}
	expectNode2 := ListNode{0, &expectNode1}
	expectHeadNode := ListNode{7, &expectNode2}
	assert.Equal(b, &expectHeadNode, resultNode)

}

func Benchmark_Test_Add_Two_Numbers_2(t *testing.B) {
	listNodeL1 := ListNode{9, nil}
	listNodeL2 := ListNode{9, &listNodeL1}
	listNodeL3 := ListNode{9, &listNodeL2}
	listNodeL4 := ListNode{9, &listNodeL3}
	listNodeL5 := ListNode{9, &listNodeL4}
	listNodeL6 := ListNode{9, &listNodeL5}
	headNodeL := ListNode{9, &listNodeL6}

	listNodeR1 := ListNode{9, nil}
	listNodeR2 := ListNode{9, &listNodeR1}
	listNodeR3 := ListNode{9, &listNodeR2}
	headNodeR := ListNode{9, &listNodeR3}

	resultNode := AddTwoNumbers(&headNodeL, &headNodeR)

	expectNode1 := ListNode{1, nil}
	expectNode2 := ListNode{0, &expectNode1}
	expectNode3 := ListNode{0, &expectNode2}
	expectNode4 := ListNode{0, &expectNode3}
	expectNode5 := ListNode{9, &expectNode4}
	expectNode6 := ListNode{9, &expectNode5}
	expectNode7 := ListNode{9, &expectNode6}
	expectHeadNode := ListNode{8, &expectNode7}

	assert.Equal(t, &expectHeadNode, resultNode)
}
