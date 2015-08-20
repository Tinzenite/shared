package shared

import "testing"

type testEqual struct {
	one  Version
	two  Version
	want bool
}

type testIncrease struct {
	before Version
	id     string
	want   Version
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
	testIncrease := []testIncrease{
		{Version{}, "a", Version{"a": 1}},
		{Version{"a": 1}, "a", Version{"a": 2}},
		{Version{"b": 12}, "a", Version{"a": 13, "b": 12}},
		{Version{"a": 11, "b": 12}, "a", Version{"a": 13, "b": 12}},
		{Version{"c": 42, "b": 12}, "a", Version{"a": 43, "b": 12, "c": 42}}}
	for _, test := range testIncrease {
		ver := test.before
		ver.Increase(test.id)
		if !ver.Equal(test.want) {
			t.Error("Expected", test.want, "got", ver)
		}
	}
}
