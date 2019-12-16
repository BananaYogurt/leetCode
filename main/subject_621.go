package main

import (
	"fmt"
	"strconv"
)

//621任务调度器

func leastInterval(tasks []byte, n int) int {

	if n == 0 {
		return len(tasks)
	}

	taskArray := [26]int{}

	for _, v := range tasks {
		taskArray[v-'A']++
	}

	for i := 0; i < len(taskArray)-1; i++ {
		for j := i + 1; j < len(taskArray); j++ {
			if taskArray[i] < taskArray[j] {
				taskArray[i], taskArray[j] = taskArray[j], taskArray[i]
			}
		}
	}

	maxTaskCount := taskArray[0]
	maxCount := 0

	for e := range taskArray {
		if taskArray[e] == maxTaskCount {
			maxCount += 1
		}
		fmt.Println(strconv.Itoa(e) + ":" + strconv.Itoa(taskArray[e]))
	}
	fmt.Println("maxTaskCount =" + strconv.Itoa(maxTaskCount) + "; " + strconv.Itoa(maxCount))
	result := (n+1)*(maxTaskCount-1) + maxCount

	size := len(tasks)

	if result >= size {
		return result
	}
	return size
}
