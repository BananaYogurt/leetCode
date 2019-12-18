package main

import (
	"fmt"
)

/**
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

*/

type RpnStack struct {
	top  int
	data []int
}

func (stack *RpnStack) push(value int) {
	stack.top++
	stack.data[stack.top] = value
}

func (stack *RpnStack) peek() int {
	element := stack.data[stack.top]
	stack.top--
	return element
}

/*func (stack * RpnStack) isEmpty () bool {
	return stack.top == -1
}
*/
func isOperator(token string) bool {
	if token == "+" || token == "*" || token == "/" || token == "-" {
		return true
	}
	return false
}

func isNegative(token string) (bool, int) {
	tokenChar := token[:]
	i := 0
	if tokenChar[0] == '-' {
		for k := 1; k < len(token); k++ {

			i = i*10 + int(rune(tokenChar[k])) - 48
		}
		return true, -i
	}
	return false, i
}

func toInt(token string) int {
	result := 0
	negative := false
	if token[0] == '-' {
		negative = true
	}
	if !negative {
		for _, i := range token {
			result = result*10 + int(rune(i)) - 48
		}
		return result
	} else {
		l := len(token)
		for index := 1; index < l; index++ {
			result = result*10 + int(rune(token[index])) - 48
		}
		return -result
	}
}

func evalRPN(tokens []string) int {
	length := len(tokens)
	stack := RpnStack{-1, make([]int, length)}
	for _, token := range tokens {
		if !isOperator(token) {
			stack.push(toInt(token))
		} else {
			a := stack.peek()
			b := stack.peek()
			switch token {
			case "+":
				{
					//fmt.Println("操作 " + strconv.Itoa(a) + "+" + strconv.Itoa(b))
					stack.push(a + b)
				}
			case "*":
				{
					//	fmt.Println("操作 " + strconv.Itoa(a) + "*" + strconv.Itoa(b))
					stack.push(a * b)
				}
			case "/":
				//fmt.Println("操作 " + strconv.Itoa(b) + "/" + strconv.Itoa(a))
				stack.push(b / a)
			case "-":
				//fmt.Println("操作 " + strconv.Itoa(b) + "-" + strconv.Itoa(a))
				stack.push(b - a)
			}
		}
	}
	return stack.peek()
}

func main() {
	//source := []string {"2","1","+","3","*"}
	//source := []string {"4","13","5","/","+"}
	source := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println("运行结果", evalRPN(source))
}
