package week2

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMinStack_GetMin(t *testing.T) {
	Convey("Test: ", t, func() {
		obj := Constructor()

		obj.Push(4) // 4
		obj.Push(5) // 4 -> 5
		obj.Push(3) // 4 -> 5 -> 3
		obj.Push(6) // 4 -> 5 -> 3 -> 6

		So(obj.Top(), ShouldEqual, 6)
		So(obj.GetMin(), ShouldEqual, 3)

		obj.Pop()
		So(obj.Top(), ShouldEqual, 3)
		So(obj.GetMin(), ShouldEqual, 3)

		obj.Pop()
		So(obj.Top(), ShouldEqual, 5)
		So(obj.GetMin(), ShouldEqual, 4)

		obj.Pop()
		So(obj.Top(), ShouldEqual, 4)
		So(obj.GetMin(), ShouldEqual, 4)
	})
}
