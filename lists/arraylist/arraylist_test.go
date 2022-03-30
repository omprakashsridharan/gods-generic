package arraylist

import "testing"

func TestListNew(t *testing.T) {
	list1 := New[int]()

	if av := list1.Empty(); av != true {
		t.Errorf("Got %v expected %v", av, true)
	}

	list2 := New(1, 2, 3)

	if av := list2.Size(); av != 3 {
		t.Errorf("Got %v expected %v", av, 3)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v %v", actualValue, 3, ok)
	}

	if _, ok := list2.Get(2); !ok {
		t.Errorf("Got %v expected %v", ok, false)
	}
}

func TestListAdd(t *testing.T) {
	list := New[string]()
	list.Add("hello", "world")
	list.Add("bye")
	if av := list.Empty(); av != false {
		t.Errorf("Got %v expected %v", av, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "bye" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "bye")
	}
}
