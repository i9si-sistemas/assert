package assert

import (
	"slices"
	"testing"
)

func TestLength(t *testing.T) {
	spy := newSpy(t)
	fatalMessage := "mySlice should have length 3"
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4)
	Length(spy, mySlice, 3, fatalMessage)
	if !slices.Equal(spy.params["Fatal"].Args, []any{fatalMessage}) {
		t.Fail()
	}
	if !spy.calls["Fatal"] {
		t.Fatal("not called")
	}
}
