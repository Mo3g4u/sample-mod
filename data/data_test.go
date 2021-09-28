package data

import (
	"testing"
)

func TestData(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		r := SetData("aaa", "aaa")

		if r["aaa"] != "aaa" {
			t.Fatal("error")
		}

	})

}
