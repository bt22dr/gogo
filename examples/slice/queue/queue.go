package queue

import (
	"fmt"
	"strconv"
)

func QueueSample(seq []string) {
	var q = make([]int, 0, 10)
	for i := 0; i < len(seq); i++ {
		if len(q) > 0 && seq[i] == "pop" {
			q = q[1:]
		} else if n, err := strconv.Atoi(seq[i]); err == nil {
			q = append(q, n)
		}
		fmt.Println(q)
	}
}

func MakeMyQueue(size int) []int {
	var q = make([]int, 0, size)
	return q
}

func Insert(q []int, num int) []int {
	q = append(q, num)
	return q
}

func Pop(q []int) []int {
	if len(q) > 0 {
		q = q[1:]
	}
	return q
}
