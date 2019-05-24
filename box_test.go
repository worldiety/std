package std

import "testing"

func TestBox_SetInt(t *testing.T) {
	box := &Box{}
	box.SetInt64(1234)
	if !box.IsInt64() {
		t.Fatal("expected int64")
	}
	if box.Int64() != 1234 {
		t.Fatal("expected", 1234)
	}

	box.SetInt(2345)
	if box.Int() != 2345 {
		t.Fatal("expected", 2345)
	}
}
