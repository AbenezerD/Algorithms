package hard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFullJustify(t *testing.T) {
	Convey("Test: fullJustify", t, func() {
		So(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16), ShouldResemble,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		)
		So(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16), ShouldResemble,
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		)
		So(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20), ShouldResemble,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		)
	})
}
