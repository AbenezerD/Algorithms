package main

import (
	"testing"
)

func TestDay25(t *testing.T) {
	Convey("TestDay25", t, func() {
		//encryptionKey, err := day25(5764801, 17807724)
		//So(err, ShouldBeNil)
		//So(encryptionKey, ShouldEqual, 14897079)

		encryptionKey, err := day25(16616892, 14505727)
		So(err, ShouldBeNil)
		So(encryptionKey, ShouldEqual, 4441893)
	})
}
