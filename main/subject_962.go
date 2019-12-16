package main

import "fmt"

//给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
//
//找出 A 中的坡的最大宽度，如果不存在，返回 0 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-width-ramp
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxWidthRamp(A []int) int {

	length := len(A)
	for k := length - 1; k > 0; k-- { // k是有可能的最大长度
		for i := 0; i+k < length; i++ {

			if A[k+i] >= A[i] {
				return k
			}
		}
	}

	return 0
}

func main() {
	a := [...]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	//a := [...]int {1, 0}
	s := a[:]
	fmt.Println("result ", maxWidthRamp(s))
}
