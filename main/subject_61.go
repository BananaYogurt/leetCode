package main

import "fmt"

/**
61. 旋转链表
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	offset := k % length
	fmt.Println("v1, v2", length, offset)
	if offset == 0 || length == 1 {
		return head
	}
	index := 0
	p1 := head
	var newHead *ListNode
	for p1 != nil {
		index++
		if index == length-offset {
			newHead = p1.Next
			p1.Next = nil
			p1 = newHead
			index++
		}
		if index == length {
			p1.Next = head
			break
		}
		p1 = p1.Next
	}

	return newHead

}
