package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Value int
	Next  *ListNode
}

var root *ListNode = nil

func Enqueue(t *ListNode, v int) *ListNode {
	if t == nil {
		root = &ListNode{Value: v, Next: nil}
		return root
	}

	if t.Value == v {
		fmt.Printf("ListNode already exists: %d\n", v)
		return root
	}

	if t.Next == nil {
		t.Next = &ListNode{Value: v, Next: nil}
		return root
	}

	return Enqueue(t.Next, v)
}

func GetHeadPtr() *ListNode {
	return root
}

func Traverse(t *ListNode) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	for t != nil {
		fmt.Printf("-> %d ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

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

func main() {
	root := GetHeadPtr()

	for i := 0; i < 6; i++ {
		root = Enqueue(root, i+1)
	}
	Traverse(root)

	temp := root
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}

	sliceString := make([]int, 0)
	mid := count / 2
	temp = root
	tempCount := 0
	for temp != nil {
		if tempCount >= mid {
			sliceString = append(sliceString, temp.Value)
		}
		temp = temp.Next
		tempCount++
	}

	fmt.Println(count, mid)

	fmt.Println(sliceString)

}
