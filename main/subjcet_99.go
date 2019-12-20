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
	nodes [8]*TmpNode
}

type TmpNode struct {
	node     *TreeNode
	nodeType int // 0 根 ； 1 右边 ； -1 左边
}

func (stack *TmpStack) push(node *TmpNode) {
	stack.top++
	stack.nodes[stack.top] = node
}

func (stack *TmpStack) peek() *TmpNode {
	node := stack.nodes[stack.top]
	stack.top--
	return node
}

func (tmpNode *TmpNode) getLeft() *TreeNode {
	return tmpNode.node.Left
}

func (tmpNode *TmpNode) getRight() *TreeNode {
	return tmpNode.node.Right
}

func recoverTree(root *TreeNode) {
	var n1 *TreeNode
	var n2 *TreeNode
	stack := TmpStack{-1, [8]*TmpNode{}}
	if root != nil {
		rootV := root.Val
		deep := -1
		stack.push(&TmpNode{root, 0})
		for stack.top > -1 {
			curNode := stack.peek()
			left := curNode.getLeft()
			right := curNode.getRight()
			ok := true
			deep++
			if nil != right {
				if curNode.nodeType == 0 {
					stack.push(&TmpNode{right, 1})
				} else {
					stack.push(&TmpNode{right, curNode.nodeType})
				}
				if right.Val < curNode.node.Val {
					ok = false
				}
			}
			if nil != left {
				if curNode.nodeType == 0 {
					stack.push(&TmpNode{left, -1})
				} else {
					stack.push(&TmpNode{left, curNode.nodeType})
				}
				if left.Val > curNode.node.Val {
					ok = false
				}
			}
			if curNode.nodeType == -1 { //左子树的节点值全部小于根结点的值
				if curNode.node.Val > rootV {
					if deep == 1 {
						n1 = root.Left
						if root.Right == nil {
							n2 = root
						} else {
							n2 = root.Right
						}
					} else {
						n1 = root
						n2 = curNode.node
					}
					break
				}
			} else if curNode.nodeType == 1 { //右子树的节点值全部大于根结点的值
				if curNode.node.Val < rootV {
					if deep == 1 {
						n1 = root.Right
						if root.Left == nil {
							n2 = root
						} else {
							n2 = root.Left
						}
					} else {
						n1 = root
						n2 = curNode.node
					}

					break
				}
			}

			if !ok {
				if n1 == nil {
					n1 = curNode.node
				} else {
					n2 = curNode.node
					break
				}
			}
		}
		if n1 != nil && n2 != nil {
			n1.Val, n2.Val = n2.Val, n1.Val
		}
	}
}

/*func main() {
	node6 := TreeNode{2, nil, nil}
	node3 := TreeNode{4, &node6, nil}
	node2 := TreeNode{1, nil, nil}
	node1 := TreeNode{3, &node2, &node3}
	fmt.Println("恢复前")
	printNode(&node1)
	recoverTree(&node1)
	fmt.Println("恢复后")
	printNode(&node1)
}
*/
func printNode(node *TreeNode) {
	/*prefix := "["
	suffix := "]"*/
	if node != nil {
		fmt.Print(node.Val)
		fmt.Print("\t")
		printNode(node.Left)
		printNode(node.Right)
	} else {
		fmt.Print("null\t")
	}
}
