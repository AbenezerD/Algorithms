package medium

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterCombinations(t *testing.T) {
	Convey("Test: letterCombinations", t, func() {
		So(letterCombinations(""), ShouldBeEmpty)
		So(letterCombinations("2"), ShouldResemble, []string{"a", "b", "c"})
		So(letterCombinations("23"), ShouldResemble, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	})
}
