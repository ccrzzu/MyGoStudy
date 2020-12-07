package array

import (
	"fmt"
	"sort"
	"strings"
)

//加一操作
func plusOne(digits []int) []int {
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- {
		//最后一位肯定要加1
		if i == len(digits)-1 {
			digits[i]++
		} else {
			//判断非最后一位要不要+1
			digits[i] += addon
			addon = 0
		}
		//非最后一位，如果等于10，要+1
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	if addon == 1 {
		result = []int{1}
		result = append(result, digits...)
	} else {
		result = digits
	}
	return result
}

//两个数组的交集，结果中是需要包含重复次数的
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item]++
	}
	for _, item := range nums2 {
		if mp1[item] > 0 {
			mp1[item]--
			res = append(res, item)
		}
	}
	return res
}

//让两个数组排序后，然后用双指针策略
func intersectBySort(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return res
}

//让两个数组排序后，然后用双指针策略
func intersectBySortNoNewMem(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums1[k] = nums2[j]
			i++
			j++
			k++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		}
	}
	return nums1[:k]
}

//两个数组的交集，结果中重复次数只返回一次
func intersectUnique(nums1 []int, nums2 []int) []int {
	var res []int
	mp1 := make(map[int]int, 0)
	for _, item := range nums1 {
		mp1[item] = 0
	}
	for _, item := range nums2 {
		if _, ok := mp1[item]; ok {
			res = append(res, item)
			delete(mp1, item)
		}
	}
	return res
}

//原地删除
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//删除有序数组的重复元素
func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//移动0
func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			k++
		} else {
			i++
		}
	}
	for i := 0; i < k; i++ {
		nums = append(nums, 0)
	}
}

//字符串数组的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, item := range strs {
		for strings.Index(item, prefix) != 0 {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

//合并两个排序的数组，合并B到A，形成一个新的有序数组
func Merge(A []int, m int, B []int, n int) {
	var res []int
	var i, j int
	for i < m && j < n {
		if A[i] <= B[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
		fmt.Println(res)
	}
	if i < m {
		res = append(res, A[i:m]...)
	}
	if j < n {
		res = append(res, B[j:n]...)
	}
	for i, item := range res {
		A[i] = item
	}
}

//归并排序逻辑
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, tail := m-1, n-1, m+n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[tail] = nums1[i1]
			i1--
		} else {
			nums1[tail] = nums2[i2]
			i2--
		}
		tail--
	}
	for tail >= 0 && i1 >= 0 {
		nums1[tail] = nums1[i1]
		i1--
		tail--
	}
	for tail >= 0 && i2 >= 0 {
		nums1[tail] = nums2[i2]
		i2--
		tail--
	}
}
