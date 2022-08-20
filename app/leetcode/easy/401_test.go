package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadBinaryWatch(t *testing.T) {
	Convey("Test: readBinaryWatch", t, func() {
		So(readBinaryWatch(1), ShouldResemble,
			[]string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"})
		So(readBinaryWatch(9), ShouldResemble,
			[]string{})
	})
}
