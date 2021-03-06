package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	max := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func mySqrt(x int) int {
	guess := 1.0 //math.MaxFloat64
	cnt := 0
	for !(math.Abs(guess*guess-float64(x)) < 1) {
		// fmt.Printf("%f", guess)
		guess = guess/2.0 + float64(x)/(2*guess)
		cnt++
	}
	// fmt.Printf("%d for %d\n", cnt, x)
	return int(guess)
}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	i := 2
	for {
		t := n2
		n2 = n1 + n2
		n1 = t
		i++
		if i == n {
			return n2
		}
	}
}
