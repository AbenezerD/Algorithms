package problems

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateArray(t *testing.T) {
	Convey("TestRotateArray", t, func() {
		So(rotate([]int{1,2,3,4,5,6,7}, 3), ShouldResemble, []int{5,6,7,1,2,3,4})
	})
}
