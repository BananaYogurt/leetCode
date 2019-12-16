package main

/**
接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
func Trap(height []int) int {
	//fmt.Println(height)
	if len(height) < 2 {
		return 0
	}

	left := 0
	leftP := -1
	rightP := -1
	capacity := 0 //最终容量
	b := false
	length := len(height)
	for i := 0; i < length; i++ {
		h := height[i]
		if left == 0 {
			if h == 0 {
				leftP++
				continue
			} else {
				left = height[i]
				leftP = i
				continue
			}
		}
		if h >= left {
			rightP = i
			//开始计算
			for k := leftP + 1; k < rightP; k++ {
				capacity += left - height[k]
			}
			leftP = i
			left = h
			b = true
		} else {
			b = false
		}
	}
	if b {
		return capacity
	} else {
		newLen := length - (leftP)
		newArray := make([]int, newLen)
		for j := 0; j < newLen; j++ {
			newArray[j] = height[length-1-j]
		}
		return capacity + Trap(newArray)
	}
}
