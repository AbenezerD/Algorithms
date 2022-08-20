package hard

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsNumber(t *testing.T) {
	Convey("Test: isNumber - valid", t, func() {
		valid := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "46.e3",
			"2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}

		for _, s := range valid {
			So(fmt.Sprintf("isNumber(%s) = %s", s, strconv.FormatBool(isNumber(s))), ShouldEqual,
				fmt.Sprintf("isNumber(%s) = true", s))
		}
	})

	Convey("Test: isNumber - invalid", t, func() {
		invalid := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", ".", "+e3"}

		for _, s := range invalid {
			So(fmt.Sprintf("isNumber(%s) = %s", s, strconv.FormatBool(isNumber(s))), ShouldEqual,
				fmt.Sprintf("isNumber(%s) = false", s))
		}
	})
}
