package week4

func canJump(nums []int) bool {
	visited := make(map[int]struct{}, 0)
	currentIndex := 0
	for i := currentIndex; i < nums[currentIndex]; i++ {

		if currentIndex >= len(nums)-1 {
			return true
		}
		if _, ok := visited[currentIndex]; ok {
			return false
		}
		visited[currentIndex] = struct{}{}
	}
}
