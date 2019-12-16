package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMaxWidthRamp(t *testing.T) {
	t1 := time.Now().Unix()
	array := [...]int{15942, 40892, 41699, 45505, 27967, 40539, 49295, 35491, 48517, 23679, 49477, 37060, 43788, 31154, 28806, 4935, 14039, 15785, 22669}
	s := array[:]
	out := maxWidthRamp(s)
	t2 := time.Now().Unix()
	fmt.Println("962结果是", out, t2-t1)
}
