package easy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPrime(t *testing.T) {
	Convey("Test: isPrime", t, func() {
		So(isPrime(13), ShouldBeTrue)
		So(isPrime(11), ShouldBeTrue)
		So(isPrime(15), ShouldBeFalse)
		So(isPrime(2), ShouldBeTrue)
		So(isPrime(123), ShouldBeFalse)
		So(isPrime(10), ShouldBeFalse)

		arr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
		for _, i := range arr {
			So(isPrime(i), ShouldBeTrue)
		}
	})

	Convey("Test: countPrimes", t, func() {
		So(countPrimes(10), ShouldEqual, 4)
		So(countPrimes(0), ShouldEqual, 0)
		So(countPrimes(1), ShouldEqual, 0)
	})
}
