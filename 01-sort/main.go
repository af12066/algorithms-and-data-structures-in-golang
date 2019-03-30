package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

var sort = make([]int, N)

func BubbleSort() {
	var flag bool
	for {
		flag = false
		for i := 0; i < N-1; i++ {
			if sort[i] > sort[i+1] {
				flag = true
				j := sort[i+1]
				sort[i+1] = sort[i]
				sort[i] = j
			}
		}
		if !flag {
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		sort[i] = rand.Intn(N * 100)
		fmt.Printf("Number %d: %d\n", i, sort[i])
	}
	fmt.Println(time.Now())
	BubbleSort()
	fmt.Println(time.Now())
	for i := 0; i < N; i++ {
		fmt.Printf("Number %d: %d\n", i, sort[i])
	}
}
