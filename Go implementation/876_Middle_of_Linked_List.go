package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}

	mid := count / 2

	temp = head
	tempCount := 0
	for temp != nil {
		if tempCount == mid {
			return temp
		}
		temp = temp.Next
		tempCount++
	}

	return nil
}
