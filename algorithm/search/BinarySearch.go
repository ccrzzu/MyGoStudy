package search

import (
	"fmt"
	"sort"
)

//二分查找
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

//二分查找所有符合的值
func binarySearchRange(nums []int, target int) []int {
	res := []int{}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) || nums[left] != target {
		res = append(res, -1)
	} else {
		res = append(res, left)
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right<0 || nums[right] != target {
		res = append(res, -1)
	} else {
		res = append(res, right)
	}
	return res
}

//go内部
func GoSearchInSort()  {
	fmt.Println(sort.SearchInts([]int{1,3,6,7,8}, 7))
}
