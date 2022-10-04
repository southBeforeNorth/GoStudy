package main

import (
	"fmt"
	"math"
)

//二分查找+插入位置
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
		if left > right {
			return left
		}
	}
}

// 递增数组 平方后 再排序
func sortedSquares(nums []int) []int {
	left := make([]int, 0)
	right := make([]int, 0)
	result := make([]int, 0)
	for _, value := range nums {
		if value < 0 {
			left = append([]int{value * value}, left...)
		} else {
			right = append(right, value*value)
		}
	}

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

//轮转数组
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

//283. 移动零
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroes(nums []int) {
	for key, value := range nums {
		if value == 0 {
			nums = append(nums, nums[0:key]...)

		}
	}
}

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

func main() {

	fmt.Println(float64(1) / float64(2))
	fmt.Println(Round(1/2, 2))
	fmt.Println("hello")
}
