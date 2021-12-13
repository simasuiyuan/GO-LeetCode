package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0} //create dummy head
	n1, n2, carry, current := 0, 0, 0, head

	/*
		golang syntax alert!

		==============================
		for loop in golang:

		for {init}; {condition}; {update condition} {
			...
		}

		e.g.
			for i := 1; i < 5; i++ {
				sum += 1
			}

		idea to note: only if {condition} is true then execute {...}
		==============================
		Infinite loop:
		e.g.
			sum := 0
			for {
				sum++
			}
			fmt.Println(sum) // never reached
		==============================
		do-while loop:

		for ok := true; ok; ok = condition {
			work()
		}

		OR

		for {
			work()
			if !condition {
				break
			}
		}


	*/

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		} //traversal first list: if no more value => n1 = 0
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		} //traversal second list: if no more value => n2 = 0
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		/*
			golang syntax alert!
			==============================
			Golang Linked List Node Struct

			type node struct {
				Val int
				Next *node
			}

			&ListNode{} give the address of the this ListNode

			* Operator: also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
			& Operator: termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

		*/
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
