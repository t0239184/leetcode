package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c_node := &ListNode{}
	head_node := c_node

	for i := 0; i < 100; i++ {
		v1 := getNodeValue(l1)
		v2 := getNodeValue(l2)
		sum, carry := calSumAndCarry(c_node.Val + v1 + v2)
		c_node.Val = sum

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			if carry > 0 {
				c_node.Next = &ListNode{carry, nil}
			}
			break
		}
		c_node.Next = &ListNode{carry, nil}
		c_node = c_node.Next
	}

	return head_node
}

func getNodeValue(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func calSumAndCarry(value int) (sum, carry int) {
	if value > 9 {
		return value % 10, 1
	}
	return value, 0 
}
