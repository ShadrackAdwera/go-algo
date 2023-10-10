package warmup

import "math"

func maxSubSliceSumNaive(nums []int, lenSubSlice int) int {
	if len(nums) < lenSubSlice {
		return 0
	}
	max := math.MinInt
	for i := 0; i < len(nums)-lenSubSlice+1; i++ {
		temp := 0
		for j := 0; j < lenSubSlice; j++ {
			temp += nums[i+j]
		}
		if temp > max {
			max = temp
		}
	}
	return max
}

func maxSubSliceSumOptimized(nums []int, lenSubSlice int) int {
	if len(nums) < lenSubSlice {
		return 0
	}

	// Calculate initial window sum
	windowSum := 0
	for i := 0; i < lenSubSlice; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum

	// Sliding window approach
	for i := lenSubSlice; i < len(nums); i++ {
		windowSum = windowSum - nums[i-lenSubSlice] + nums[i]
		maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
	}

	return maxSum
}
