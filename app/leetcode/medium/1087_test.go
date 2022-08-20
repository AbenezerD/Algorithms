package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExpand(t *testing.T) {
	Convey("Test: expand", t, func() {
		So(expand("abcd"), ShouldResemble, []string{"abcd"})
		So(expand("a{b,c}de{f,g}"), ShouldResemble, []string{"abdef", "abdeg", "acdef", "acdeg"})
		So(expand("{a,b}c{d,e}f"), ShouldResemble, []string{"acdf", "acef", "bcdf", "bcef"})
		So(expand("{f,g,h}i}"), ShouldResemble, []string{"fi", "gi", "hi"})
		So(expand("{a}{b}"), ShouldResemble, []string{"ab"})
		So(expand("{a,b}{z,x,y}"), ShouldResemble, []string{"ax", "ay", "az", "bx", "by", "bz"})
	})
}
