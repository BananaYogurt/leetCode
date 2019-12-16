package main

import (
	"math/rand"
	"time"
)

//leetcode 382 链表随机节点

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	//rand.Int31()
	res := this.head.Val
	i := 2
	tmp := this.head.Next
	now := time.Now().UnixNano()
	rand.Seed(now)
	for tmp != nil {
		randNum := rand.Int()
		if randNum%i == 0 {
			res = tmp.Val
		}
		i++
		tmp = tmp.Next
	}
	return res
}

/*func main() {
	p5 := ListNode{5, nil}
	p4 := ListNode{4, &p5}
	p3 := ListNode{3, &p4}
	p2 := ListNode{2, &p3}
	p1 := ListNode{1,  &p2}
	p0 := ListNode{0, &p1 }
	obj := Constructor(&p0)
	r := obj.GetRandom()
	fmt.Println("随机结果 = " , r)
}*/
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
