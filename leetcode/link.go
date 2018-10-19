package main

import "fmt"

func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	//fmt.Println(removeNthFromEnd(&node5, 1))
	//node1 := ListNode{1,
	//	&ListNode{2,
	//	&ListNode{4, nil}}}
	//fmt.Println(reverseList(&node1))
	fmt.Println(middleNode(&node1))
	//fmt.Println(mergeTwoLists(&node1, &node2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * 19.给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var next = head.Next
	for i := 1; i < n; i++ {
		next = next.Next
	}

	if nil == next {
		return head.Next
	}

	var mark = head
	for {
		next = next.Next
		if nil == next {
			break
		}
		mark = mark.Next
	}

	mark.Next = mark.Next.Next
	return head
}

/**
 * 21. 将两个有序链表合并为一个新的有序链表并返回。
 * 新链表是通过拼接给定的两个链表的所有节点组成的。
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	var head *ListNode
	if l1.Val < l2.Val {
		head = l1

		l1 = l1.Next
		if nil == l1 {
			head.Next = l2
			return head
		}
	} else {
		head = l2

		l2 = l2.Next
		if nil == l2 {
			head.Next = l1
			return head
		}
	}

	var list = head

	for {
		if l1.Val < l2.Val {
			list.Next = l1
			l1 = l1.Next
			if nil == l1 {
				list.Next.Next = l2
				break
			}
		} else {
			list.Next = l2
			l2 = l2.Next
			if nil == l2 {
				list.Next.Next = l1
				break
			}
		}
		list = list.Next
	}

	return head
}

/**
 * 206. 反转一个单链表。
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	stack := make([]*ListNode, 0, 8)

	for nil != head {
		stack = append(stack, head)
		head = head.Next
	}

	length := len(stack)
	list := stack[length-1]
	for i:=length-2; i>-1; i-- {
		list.Next = stack[i]
		list = list.Next
	}
	list.Next = nil

	return stack[length-1]
}

/**
 * 876. 给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
 * 如果有两个中间结点，则返回第二个中间结点。
 */
func middleNode(head *ListNode) *ListNode {
	if nil == head {
		return head
	}

	slow := head
	fast := head.Next
	if nil == fast {
		return slow
	}

	for {
		slow = slow.Next
		fast = fast.Next
		if nil == fast {
			break
		}
		fast = fast.Next
		if nil == fast {
			break
		}
	}
	return slow
}