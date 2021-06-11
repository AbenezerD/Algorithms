package problems

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestPalindrome(t *testing.T) {
	Convey("Test: LongestPalindrome", t, func() {
		So(longestPalindrome("babad"), ShouldBeIn, []string{"bab", "aba"})
		So(longestPalindrome("cbbd"), ShouldEqual, "bb")
		So(longestPalindrome("a"), ShouldEqual, "a")
		So(longestPalindrome("ac"), ShouldBeIn, []string{"a", "c"})
	})
}

func TestIsPalindrome(t *testing.T) {
	Convey("Test: IsPalindrome", t, func() {
		So(isPalindrome("a"), ShouldBeTrue)
		So(isPalindrome("abba"), ShouldBeTrue)
		So(isPalindrome("aba"), ShouldBeTrue)
		So(isPalindrome("abcdedcba"), ShouldBeTrue)
		So(isPalindrome("abcdedcba"), ShouldBeTrue)
		So(isPalindrome("abcdefdcba"), ShouldBeFalse)
	})
}
