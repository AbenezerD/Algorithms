package week1

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	Convey("Given: ", t, func() {
		So(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), ShouldResemble, [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		})
	})
}
