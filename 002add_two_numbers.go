package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

CONSTRAINT

    The number of nodes in each linked list is in the range [1, 100].
    0 <= Node.val <= 9
    It is guaranteed that the list represents a number that does not have leading zeros.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result  = &ListNode{}
		curNode = result
		rem     int
	)

	for {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		val := l1Val + l2Val + rem
		if val >= 10 {
			val = val - 10
			rem = 1
		} else {
			rem = 0
		}

		curNode.Val = val

		if l1 == nil && l2 == nil && rem == 0 {
			break
		}

		curNode.Next = &ListNode{}
		curNode = curNode.Next
	}

	return result
}

func add(l1 *ListNode, l2 *ListNode) *ListNode {
	currentNode := &ListNode{}
	result := currentNode
	rem := 0

	for {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		v := val1 + val2 + rem
		if v >= 10 {
			v = v - 10
			rem = 1
		} else {
			rem = 0
		}
		currentNode.Val = v
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil && rem == 0 {
			break
		}
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}
	return result
}

func main() {
	a := &ListNode{8, &ListNode{2, nil}}
	b := &ListNode{4, &ListNode{7, &ListNode{6, nil}}}
	result := add(a, b)
	for {
		fmt.Println(result.Val)
		result = result.Next
		if result == nil {
			break
		}
	}
}

//func test(ptr *int) {
//	*ptr = *ptr + 1
//	fmt.Println("test : ", *ptr)
//}
//
//func main() {
//	num := 10
//	ptr := &num
//	test(ptr)
//	fmt.Println("main : ", num)
//}
