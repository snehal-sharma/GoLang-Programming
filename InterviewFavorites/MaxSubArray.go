// Also known as Kadane's Algorithm
package main

import (
	"fmt"
)

func MaxSubArraySum(arr []int) (int, int, int) {

	maxSum := 0
	currentSum := 0
	start, end := 0, 0
	tempStart := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]+currentSum > arr[i] {
			currentSum += arr[i]
		} else {
			currentSum = arr[i]
			tempStart = i

		}

		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}
	fmt.Println(start, end, maxSum)
	return start, end, maxSum
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	start, end, max := MaxSubArraySum(arr)
	fmt.Printf("Maximum subarray sum is %d Start:%d End:%d", max, start, end)
}
