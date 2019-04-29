package std

import "testing"

func Test_list_Add(t *testing.T) {
	list := &StrList{}
	list.Add("hello")
	if list.Get(0) != "hello" {
		t.Fatal("unexpected", list.Get(0))
	}
}
