package week4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLRUCache(t *testing.T) {
	Convey("Test: LRUCache", t, func() {
		cache := Constructor(3 /* capacity */)

		cache.Put(1, 1)
		cache.Put(2, 2)
		So(cache.Get(1), ShouldEqual, 1)  // returns 1
		cache.Put(3, 3)                   //
		So(cache.Get(2), ShouldEqual, 2)  // returns -1 (not found)
		cache.Put(4, 4)                   // evicts key 2
		So(cache.Get(1), ShouldEqual, -1) // returns -1 (not found)
		So(cache.Get(3), ShouldEqual, 3)  // returns 3
		So(cache.Get(4), ShouldEqual, 4)  // returns 4
		cache.Put(5, 5)                   // evicts key 2
		cache.Put(6, 6)                   // evicts key 3
		cache.Put(7, 7)                   // evicts key 1
		So(cache.Get(1), ShouldEqual, -1) // returns -1 (not found)
		So(cache.Get(3), ShouldEqual, -1) // returns 3
		So(cache.Get(4), ShouldEqual, -1) // returns 4
		So(cache.Get(5), ShouldEqual, 5)  // returns -1 (not found)
		So(cache.Get(6), ShouldEqual, 6)  // returns 3
		So(cache.Get(7), ShouldEqual, 7)  // returns 4

	})
}
