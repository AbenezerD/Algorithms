package medium

/**
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:
1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/

func dailyTemperatures(temperatures []int) []int {
	diff := make([]int, len(temperatures))
	for i := 0; i < len(temperatures)-1; i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				diff[i] = j - i
				break
			}
		}
	}
	return diff
}

// [73,74,75,71,69,72,76,73]
// [0, +1, +1, -4, -2, +3, +4, -3]
// [0, +1, +2, -2, -6, -3, +1, -2]
// [0, +1, +2, -2, -6, -3, +1, -2]
// [1,1,4,2,1,1,0,0]
