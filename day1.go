package main

import (
	"math"
	"sort"
)

// 判断一个数是否是质数
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    sqrtN := int(math.Sqrt(float64(n)))
    for i := 2; i <= sqrtN; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func Prime(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func Deduplicate(nums []int) []int{
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}
	result := []int{}
	for num := range unique {
		result = append(result, num)
	}
	return result
}

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func BuildTree(nums []int, i int) *TreeNode {
	if i >= len(nums) || nums[i] == -1 {
		return nil
	}
	root := &TreeNode{Value: nums[i]}
	root.Left = BuildTree(nums, 2*i+1)
	root.Right = BuildTree(nums, 2*i+2)
	return root
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
    nums3 := make([]int, m+n)
    for i := 0; i < m; i++ {
        nums3[i] = nums1[i]
    }
    for i := 0; i < n; i++ {
        nums3[m+i] = nums2[i]
    }
    sort.Ints(nums3)
    return nums3
}
