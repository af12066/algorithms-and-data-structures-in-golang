package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 5

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

func QuickSort(data []int, left, right int) {
	var lower_idx, upper_idx int
	if left >= right {
		return
	}
	// とりあえず左端の値をピボットとする
	var pivot int = data[left]
	for lower_idx, upper_idx = left, right; lower_idx < upper_idx; {
		/* 右端から中心に向かって、それぞれの値がピボットより大きいかどうか比較する。
		   ピボットより小さい場合は、以下のif statementにより、左側のピボットより大きい値と入れ替える。
		*/
		for data[upper_idx] > pivot && lower_idx < upper_idx {
			upper_idx--
		}
		/* 左端から中心に向かって、それぞれの値がピボットより小さいかどうか比較する。
		   ピボットより大きい場合は、以下のif statementにより、右側のピボットより小さい値と入れ替える。
		*/
		for data[lower_idx] <= pivot && lower_idx < upper_idx {
			lower_idx++
		}
		if lower_idx < upper_idx {
			tmp := data[lower_idx]
			data[lower_idx] = data[upper_idx]
			data[upper_idx] = tmp
		}
	}
	// ピボットを更新する
	tmp := data[left]
	data[left] = data[upper_idx]
	data[upper_idx] = tmp
	// 上記で決定したupper_idxより左側について同様にソート
	// 区間を短くすることで計算量が少なく済む
	QuickSort(data, left, upper_idx-1)
	// 上記で決定したupper_idxより右側について同様にソート
	QuickSort(data, upper_idx+1, right)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		sort[i] = rand.Intn(N * 100)
		fmt.Printf("Number %d: %d\n", i, sort[i])
	}
	fmt.Println(time.Now())
	// BubbleSort()
	QuickSort(sort, 0, N-1) // まずはsliceの先頭から末尾まで比較する
	fmt.Println(time.Now())
	for i := 0; i < N; i++ {
		fmt.Printf("Number %d: %d\n", i, sort[i])
	}
}
