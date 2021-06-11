package main

import (
	"testing"
)

func TestDay23(t *testing.T) {
	Convey("TestDay23", t, func() {
		So(day23c("389125467", 10000000, 1000000), ShouldEqual, "149245887792") // 1001
		So(day23c("219347865", 10000000, 1000000), ShouldEqual, "149245887792") // 1001
	})
}
