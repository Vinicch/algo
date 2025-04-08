package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(binarySearch([]int{2, 3, 4, 5, 6, 7, 8, 9, 10}, 99))

	arr := []int{10, 5, 9, 6, 2, 7, 4, 8, 1, 3}
	bubbleSort(arr)
	fmt.Println(arr)

	arr = []int{10, 5, 9, 6, 2, 7, 4, 8, 1, 3}
	insertionSort(arr)
	fmt.Println(arr)

	arr = []int{10, 5, 9, 6, 2, 7, 4, 8, 1, 3}
	quickSort(arr)
	fmt.Println(arr)

	arr = []int{10, 5, 9, 6, 2, 7, 4, 8, 1, 3}
	fmt.Println(windowSum(arr, 3))
}

func binarySearch(nums []int, target int) bool {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return false
}

func bubbleSort(nums []int) {
	for i := range nums {
		for j := range len(nums) - 1 - i {
			if nums[j] > nums[j+1] {
				nums[j] ^= nums[j+1]
				nums[j+1] ^= nums[j]
				nums[j] ^= nums[j+1]
			}
		}
	}
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j] ^= nums[j-1]
			nums[j-1] ^= nums[j]
			nums[j] ^= nums[j-1]
		}
	}
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := partition(arr, lo, hi)

	qs(arr, lo, pivotIndex-1)
	qs(arr, pivotIndex+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	pivotIndex := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			pivotIndex++
			tmp := arr[i]
			arr[i] = arr[pivotIndex]
			arr[pivotIndex] = tmp
		}
	}

	pivotIndex++
	arr[hi] = arr[pivotIndex]
	arr[pivotIndex] = pivot

	return pivotIndex
}

func windowSum(arr []int, k int) int {
	currSum := 0
	for i := range k {
		currSum += arr[i]
	}

	maxSum := currSum
	for i := k; i < len(arr); i++ {
		currSum = currSum - arr[i-k] + arr[i]
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
