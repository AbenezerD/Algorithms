package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestSubstring(t *testing.T) {
	Convey("Test: LongestSubstring", t, func() {
		So(lengthOfLongestSubstring("abcabcab"), ShouldEqual, 3)
		So(lengthOfLongestSubstring("bbbbbb"), ShouldEqual, 1)
		So(lengthOfLongestSubstring("pwwkew"), ShouldEqual, 3)
		So(lengthOfLongestSubstring(""), ShouldEqual, 0)
		So(lengthOfLongestSubstring("abcabcabd"), ShouldEqual, 4)
		So(lengthOfLongestSubstring("bacabcabd"), ShouldEqual, 4)
		So(lengthOfLongestSubstring("bacabcadbebd"), ShouldEqual, 5)
		So(lengthOfLongestSubstring("bacabfbcadbebd"), ShouldEqual, 5)
		So(lengthOfLongestSubstring("a"), ShouldEqual, 1)
		So(lengthOfLongestSubstring("abcdef"), ShouldEqual, 6)
	})
}
