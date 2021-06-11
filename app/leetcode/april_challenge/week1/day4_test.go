package week1

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMoveZeroes(t *testing.T) {
	Convey("Given: ", t, func() {
		So(moveZeroes([]int{0,1,0,3,12}), ShouldResemble, []int{1,3,12,0,0})
		So(moveZeroes([]int{1,0,3,0,12,5,0,9}), ShouldResemble, []int{1,3,12,5,9,0,0,0})
	})
}
