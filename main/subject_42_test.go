package main

import (
	"fmt"
	"testing"
)

//42题
func TestTrap(t *testing.T) {

	array := [...]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1} //测试数组
	//array  := [...]int {8,5,4,1,8,9,3,0,0} //测试数组
	// array  := [...]int {4,2, 3} //测试数组
	testSlice := array[:]
	capacity := Trap(testSlice)
	fmt.Println("42题接雨水的最终结果为", capacity)

}
