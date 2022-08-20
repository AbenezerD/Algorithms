package hard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBraceExpansionII(t *testing.T) {
	Convey("Test: braceExpansionII", t, func() {
		So(_braceExpansionII("{a{x,ia,o}w,{n,{g,{u,o}},{a,{x,ia,o},w}},er}"), ShouldResemble,
			[]string{"a", "aiaw", "aow", "axw", "er", "g", "ia", "n", "o", "u", "w", "x"})
		So(_braceExpansionII("{a{x,ia,o}w,{n,{g,{u,o}},{a,{x,ia,o},w}},er}"), ShouldResemble,
			[]string{"a", "aiaw", "aow", "axw", "er", "g", "ia", "n", "o", "u", "w", "x"})
		//So(braceExpansionII("{{a,l},ex,{ab,c}}"), ShouldResemble, []string{"a", "ab", "c", "ex", "l"})
		//So(braceExpansionII("{{a,b},b,c}"), ShouldResemble, []string{"a", "b", "c"})

		//So(braceExpansionII("abcd"), ShouldResemble, []string{"abcd"})
		So(_braceExpansionII("a{b,c}de{f,g}"), ShouldResemble, []string{"abdef", "abdeg", "acdef", "acdeg"})
		So(_braceExpansionII("{a,b}c{d,e}f"), ShouldResemble, []string{"acdf", "acef", "bcdf", "bcef"})
		So(_braceExpansionII("{f,g,h}i}"), ShouldResemble, []string{"fi", "gi", "hi"})
		So(_braceExpansionII("{a}{b}"), ShouldResemble, []string{"ab"})
		So(_braceExpansionII("{a,b}{z,x,y}"), ShouldResemble, []string{"ax", "ay", "az", "bx", "by", "bz"})

		So(_braceExpansionII("{a,b}{c,{d,e}}"), ShouldResemble, []string{"ac", "ad", "ae", "bc", "bd", "be"})
		So(_braceExpansionII("{{a,z},a{b,c},{ab,z}}"), ShouldResemble, []string{"a", "ab", "ac", "z"})
		So(_braceExpansionII("n{g,{u,o}}{a,{x,ia,o},w}"), ShouldResemble,
			[]string{"nga", "ngia", "ngo", "ngw", "ngx", "noa", "noia", "noo", "now", "nox", "nua", "nuia", "nuo", "nuw", "nux"})
	})
}
