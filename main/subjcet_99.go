package main

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
		stack.push(&TmpNode{root, 0})
		for stack.top > -1 {
			curNode := stack.peek()
			left := curNode.getLeft()
			right := curNode.getRight()
			ok := true
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
					ok = false
				}
			} else if curNode.nodeType == 1 { //右子树的节点值全部大于根结点的值
				if curNode.node.Val < rootV {
					ok = false
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
	stack := TmpStack{-1, [8]*TreeNode{}}
	node := TreeNode{1, nil, nil}
	stack.push(&node)
	n1 := stack.peek()
	fmt.Println("运行结果", n1.Val)
}*/
