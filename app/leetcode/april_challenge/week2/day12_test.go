package week2

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLastStoneWeight(t *testing.T) {
	Convey("Test: LastStoneWeight ", t, func() {
		So(lastStoneWeight([]int{2,7,4,1,8,1}), ShouldEqual, 1)
	})
}

//				 	1
//			2				1
//     7		8		4


//				 	8
//			7				4
//     1		2		1

