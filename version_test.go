package shared

import "testing"

func TestVersion_Equal(t *testing.T) {
	verA := Version{}
	verB := Version{}
	if !verA.Equal(verB) {
		t.Error("A should equal B")
	}
	if !verB.Equal(verA) {
		t.Error("B should equal A")
	}
}

func TestVersion_Increase(t *testing.T) {
	t.SkipNow()
	version := Version{}
	version.Increase("a")
}
