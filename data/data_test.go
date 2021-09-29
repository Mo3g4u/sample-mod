package data

import (
	"testing"
)

func TestData(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		r := CountUp()

		if r != 1 {
			t.Fatal("error")
		}
		r = CountUp()

		if r != 2 {
			t.Fatal("error")
		}

	})

}
