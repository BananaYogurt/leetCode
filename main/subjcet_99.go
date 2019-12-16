package main

import "fmt"

//99. 恢复二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TmpStack struct {
	top   int
	nodes [8]*TreeNode
}

type TmpNode struct {
	node *TreeNode
	left bool
}

func (stack *TmpStack) push(node *TreeNode) {
	stack.top++
	stack.nodes[stack.top] = node
}

func (stack *TmpStack) peek() *TreeNode {
	node := stack.nodes[stack.top]
	stack.top--
	return node
}

func recoverTree(root *TreeNode) {
	var n1 *TreeNode
	var n2 *TreeNode
	stack := TmpStack{-1, [8]*TreeNode{}}
	if root != nil {
		rootV := root.Val
		stack.push(root)
		for stack.top > -1 {
			curNode := stack.peek()
			left := curNode.Left
			right := curNode.Right
			ok := true
			if nil != right {
				stack.push(right)
				if right.Val < curNode.Val {
					ok = false
				}
			}
			if nil != left {
				if left.Val > curNode.Val {
					ok = false
				}
				stack.push(left)
			}
			if !ok {
				if n1 == nil {
					n1 = curNode
				} else {
					n2 = curNode
					break
				}
			}
		}
		if n1 != nil && n2 != nil {
			n1.Val, n2.Val = n2.Val, n1.Val
		}
	}
}

func main() {
	stack := TmpStack{-1, [8]*TreeNode{}}
	node := TreeNode{1, nil, nil}
	stack.push(&node)
	n1 := stack.peek()
	fmt.Println("运行结果", n1.Val)
}
