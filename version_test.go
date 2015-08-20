package shared

import "testing"

type testEqual struct {
	one  Version
	two  Version
	want bool
}

func TestVersion_Equal(t *testing.T) {
	testEquals := []testEqual{
		// empty
		{Version{}, Version{}, true},
		{Version{}, Version{"a": 0}, false},
		{Version{}, Version{"a": 0, "b": 12}, false},
		// single
		{Version{"a": 0}, Version{"a": 0}, true},
		{Version{"a": 0}, Version{"a": 1}, false},
		{Version{"a": 0}, Version{"b": 0}, false},
		{Version{"a": 0}, Version{"b": 1}, false},
		// multiple
		{Version{"a": 0, "b": 0}, Version{"a": 0, "b": 0}, true},
		{Version{"a": 0, "b": 0}, Version{"b": 0, "a": 0}, true},
		{Version{"a": 0, "b": 0}, Version{"a": 1, "b": 0}, false},
		{Version{"a": 0, "b": 0}, Version{"a": 0, "b": 2}, false},
		{Version{"a": 0, "b": 0}, Version{"a": 1, "b": 1}, false},
		// mixed
		{Version{"a": 0}, Version{"b": 0, "a": 0}, false},
		{Version{"a": 12}, Version{"b": 1, "a": 34}, false},
		{Version{"a": 12}, Version{"b": 1, "c": 34}, false}}
	for _, test := range testEquals {
		oneEqualTwo := test.one.Equal(test.two)
		twoEqualOne := test.two.Equal(test.one)
		if oneEqualTwo != twoEqualOne {
			t.Error("Equal not symmetrical!", test.one, test.two)
		}
		if oneEqualTwo != test.want {
			t.Error("Expected", test.want, "got", oneEqualTwo, "for", test.one, test.two)
		}
	}
}

func TestVersion_Increase(t *testing.T) {
	t.SkipNow()
	version := Version{}
	version.Increase("a")
}
