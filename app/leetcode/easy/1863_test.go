package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsetXORSum(t *testing.T) {
	Convey("Test: subsetXORSum", t, func() {
		So(subsetXORSum([]int{1, 3}), ShouldEqual, 6)
		So(subsetXORSum([]int{5, 1, 6}), ShouldEqual, 28)
		So(subsetXORSum([]int{3, 4, 5, 6, 7, 8}), ShouldEqual, 480)
	})
}
