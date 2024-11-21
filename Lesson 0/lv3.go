package main

import (
	"fmt"
)

func qs(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i := l
	j := r
	x := nums[(l+r)/2]
	for i <= j {
		for i <= r && nums[i] < x {
			i++
		}
		for j >= l && nums[j] > x {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	qs(nums, l, j)
	qs(nums, i, r)
}

func main() {
	numbers := []int{8, 4, 6, 2, 3, 6, 4}
	qs(numbers, 0, len(numbers)-1)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("最大的数是：", numbers[6])
}
