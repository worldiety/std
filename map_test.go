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

func TestMap_Put(t *testing.T) {
	mp := &Map{}
	mp.Put(NewBox("hello"), NewBox("world"))
	if !mp.Has(NewBox("hello")) {
		t.Fatal("expected hello")
	}

	if mp.Get(NewBox("hello")).Unbox() != "world" {
		t.Fatal("unexpected", mp.Get(NewBox("hello")))
	}

	mp.Delete(nil)
}
