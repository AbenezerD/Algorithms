package main

import (
	"testing"
)

func TestDay14(t *testing.T) {
	Convey("Given: ", t, func() {
		So(getBitVal(9, 0), ShouldEqual, 1) // 1001
		So(getBitVal(9, 1), ShouldEqual, 0) // 1001
		So(getBitVal(9, 2), ShouldEqual, 0) // 1001
		So(getBitVal(9, 3), ShouldEqual, 1) // 1001
		So(getBitVal(9, 4), ShouldEqual, 0) // 1001
		So(getBitVal(9, 5), ShouldEqual, 0) // 1001
		So(getBitVal(9, 6), ShouldEqual, 0) // 1001
	})
}
