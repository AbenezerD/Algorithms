package main

import (
	"testing"
)

func TestDay18(t *testing.T) {
	Convey("TestDay18", t, func() {
		//So(arithmetic("((1 + 1) * (2 * 6))"), ShouldEqual, 24)
		So(arithmetic("((1 + 1) * (2 * 6)) + ((1 + 1) * (2 * 6))"), ShouldEqual, 48)
		//So(arithmetic("2 * 3"), ShouldEqual, 6)
		So(arithmetic("(2 * 3 + (4 * 5))"), ShouldEqual, 26)
		So(arithmetic("(5 + (8 * 3 + 9 + 3 * 4 * 3))"), ShouldEqual, 437)
		So(arithmetic("(5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)))"), ShouldEqual, 12240)
		So(arithmetic("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), ShouldEqual, 13632)
		//So(arithmetic("(((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2)"+
		//	" + (((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2)"), ShouldEqual, 2*13632)
	})
}

func TestDay18B(t *testing.T) {
	Convey("TestDay18", t, func() {
		//So(arithmetic("((1 + 1) * (2 * 6))"), ShouldEqual, 24)
		So(arithmetic2("((1 + 1) * (2 * 6)) + ((1 + 1) * (2 * 6))"), ShouldEqual, 48)
		//So(arithmetic("2 * 3"), ShouldEqual, 6)
		So(arithmetic2("1 + (2 * 3) + (4 * (5 + 6))"), ShouldEqual, 51)
		So(arithmetic2("2 * 3 + (4 * 5)"), ShouldEqual, 46)
		So(arithmetic2("(5 + (8 * 3 + 9 + 3 * 4 * 3))"), ShouldEqual, 1445)
		So(arithmetic2("(5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)))"), ShouldEqual, 669060)
		So(arithmetic2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), ShouldEqual, 23340)
	})
}
