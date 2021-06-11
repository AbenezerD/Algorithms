package week1

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsHappy(t *testing.T) {
	Convey("Given: ", t, func() {
		So(isHappy(19), ShouldBeTrue)
		So(isHappy(3), ShouldBeFalse)
	})
}
