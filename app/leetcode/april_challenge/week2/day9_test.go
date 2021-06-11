package week2

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	Convey("Given: ", t, func() {
		So(backspaceCompare("ab#c", "ad#c"), ShouldBeTrue)
		So(backspaceCompare("ab##", "c#d#"), ShouldBeTrue)
		So(backspaceCompare("a##c", "#a#c"), ShouldBeTrue)
		So(backspaceCompare("a#c", "b"), ShouldBeFalse)
	})
}
