package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumberOfMatches(t *testing.T) {
	Convey("Test: numberOfMatches", t, func() {
		So(numberOfMatches(7), ShouldEqual, 6)
		So(numberOfMatches(14), ShouldEqual, 13)
		So(numberOfMatches(8), ShouldEqual, 7)
	})
}
