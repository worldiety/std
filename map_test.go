package std

import "testing"

func TestStrMap_Put(t *testing.T) {
	mp := &StrStrMap{}
	mp.Put("hello", "world")
	if !mp.Has("hello") {
		t.Fatal("expected hello")
	}

	if mp.Get("hello") != "world" {
		t.Fatal("unexpected", mp.Get("hello"))
	}
}
