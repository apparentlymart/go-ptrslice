package ptrslice

import (
	"testing"
)

func TestPointerToSlice(t *testing.T) {
	i := 5
	sl := PointerToSlice(&i).([]int)
	if len(sl) != 1 {
		t.Fatalf("wrong length %d; want 1", len(sl))
	}
	if cap(sl) != 1 {
		t.Fatalf("wrong capacity %d; want 1", cap(sl))
	}
	if sl[0] != 5 {
		t.Fatalf("wrong value %d; want 5", sl[0])
	}
	sl[0] = 2
	if *(&i) != 2 {
		t.Fatalf("value didn't change")
	}
	*(&i) = 10
	if sl[0] != 10 {
		t.Fatalf("slice didn't change")
	}
}
