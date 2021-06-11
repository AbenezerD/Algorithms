package week3

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckValidString(t *testing.T) {
	Convey("Test: CheckValidString", t, func() {
		So(checkValidString("()"), ShouldBeTrue)
		So(checkValidString("(*)"), ShouldBeTrue)
		So(checkValidString("(*()*((*())))"), ShouldBeTrue)
		So(checkValidString("(*))"), ShouldBeTrue)
		So(checkValidString("**((*"), ShouldBeFalse)
		So(checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*"), ShouldBeFalse)
		So(checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"), ShouldBeTrue)
		So(checkValidString("*((((**(**))))))((((*)))))(((**(*)))(*)"), ShouldBeTrue)
	})
}
