package main

import (
	"testing"
)

func TestDay13(t *testing.T) {
	Convey("Given: ", t, func() {
		So(getReverseModulo(40, 7), ShouldEqual, 3)
		So(getReverseModulo(56, 5), ShouldEqual, 1)
		So(getReverseModulo(35, 8), ShouldEqual, 3)
		So(getReverseModulo(166439, 12), ShouldEqual, 11)
		So(getReverseModulo(102011, 25), ShouldEqual, 16)
		So(getReverseModulo(53599, 55), ShouldEqual, 19)
		So(getReverseModulo(243257, 12), ShouldEqual, 5)
	})
}
