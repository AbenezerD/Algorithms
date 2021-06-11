package week4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCanJump(t *testing.T) {
	Convey("Test: CanJump", t, func() {
		So(canJump([]int{2, 3, 1, 1, 4}), ShouldBeTrue)
		So(canJump([]int{3, 2, 1, 0, 4}), ShouldBeFalse)
	})
}
