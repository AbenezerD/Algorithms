package week4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestCommonSubsequence(t *testing.T) {
	Convey("Test: LongestCommonSubsequence", t, func() {
		So(longestCommonSubsequence("abcde", "ace"), ShouldEqual, 3)
		So(longestCommonSubsequence("abc", "abc"), ShouldEqual, 3)
		So(longestCommonSubsequence("abc", "def"), ShouldEqual, 0)
	})
}
