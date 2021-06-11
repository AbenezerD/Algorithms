package week1

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSingleNumber(t *testing.T) {
	Convey("Given: ", t, func() {

		So(executeImplementation1([]int{2, 2, 1}), ShouldEqual, 1)
		So(executeImplementation2([]int{2, 2, 1}), ShouldEqual, 1)
		So(executeImplementation3([]int{2, 2, 1}), ShouldEqual, 1)

		So(executeImplementation1([]int{4, 1, 2, 1, 2}), ShouldEqual, 4)
		So(executeImplementation2([]int{4, 1, 2, 1, 2}), ShouldEqual, 4)
		So(executeImplementation3([]int{4, 1, 2, 1, 2}), ShouldEqual, 4)

		So(executeImplementation1([]int{4, 1, 2, 1, 2, 3, 4, 8, 8, 9, 3, 9, 11, 12, 10, 11, 12}), ShouldEqual, 10)
		So(executeImplementation2([]int{4, 1, 2, 1, 2, 3, 4, 8, 8, 9, 3, 9, 11, 12, 10, 11, 12}), ShouldEqual, 10)
		So(executeImplementation3([]int{4, 1, 2, 1, 2, 3, 4, 8, 8, 9, 3, 9, 11, 12, 10, 11, 12}), ShouldEqual, 10)
	})
}

func executeImplementation1(nums []int) int {
	fmt.Printf("Implementation 1: array: %v", nums)
	start := time.Now()
	result := singleNumber(nums)
	duration := time.Since(start)
	fmt.Printf("\tExecution time: %v\n", duration)
	return result
}

func executeImplementation2(nums []int) int {
	fmt.Printf("Implementation 2: array: %v", nums)
	start := time.Now()
	result := singleNumber2(nums)
	duration := time.Since(start)
	fmt.Printf("\tExecution time: %v\n", duration)
	return result
}

func executeImplementation3(nums []int) int {
	fmt.Printf("Implementation 3: array: %v", nums)
	start := time.Now()
	result := singleNumber3(nums)
	duration := time.Since(start)
	fmt.Printf("\tExecution time: %v\n", duration)
	return result
}
