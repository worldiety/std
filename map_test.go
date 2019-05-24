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
	mp.Put(Wrap("hello"), Wrap("world"))
	if !mp.Has(Wrap("hello")) {
		t.Fatal("expected hello")
	}

	if mp.Get(Wrap("hello")).Unbox() != "world" {
		t.Fatal("unexpected", mp.Get(Wrap("hello")))
	}

	mp.Delete(nil)
}
