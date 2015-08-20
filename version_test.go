package shared

import "testing"

type testEqual struct {
	one  Version
	two  Version
	want bool
}

type testMax struct {
	ver  Version
	want int
}

type testIncrease struct {
	before Version
	id     string
	want   Version
}

type testValid struct {
	local  Version
	selfid string
	remote Version
	want   bool
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

func TestVersion_Max(t *testing.T) {
	testMax := []testMax{
		{Version{}, 0},
		{Version{"a": 1}, 1},
		{Version{"a": 1, "b": 1}, 1},
		{Version{"a": -2, "b": 1}, 1},
		{Version{"a": 2, "b": 1}, 2},
		{Version{"a": 2, "b": 1, "c": 42}, 42},
		{Version{"a": -1}, 0}}
	for _, test := range testMax {
		max := test.ver.Max()
		if max != test.want {
			t.Error("Expected", test.want, "got", max)
		}
	}
}

func TestVersion_Increase(t *testing.T) {
	testIncreases := []testIncrease{
		{Version{}, "a", Version{"a": 1}},
		{Version{"a": 1}, "a", Version{"a": 2}},
		{Version{"b": 12}, "a", Version{"a": 13, "b": 12}},
		{Version{"a": 11, "b": 12}, "a", Version{"a": 13, "b": 12}},
		{Version{"c": 42, "b": 12}, "a", Version{"a": 43, "b": 12, "c": 42}}}
	for _, test := range testIncreases {
		ver := test.before
		ver.Increase(test.id)
		if !ver.Equal(test.want) {
			t.Error("Expected", test.want, "got", ver)
		}
	}
}

func TestVersion_Valid(t *testing.T) {
	testValids := []testValid{
		// empty OP
		{Version{}, "a", Version{}, true},
		// no OP
		{Version{"a": 1, "b": 2}, "a", Version{"a": 1, "b": 2}, true},
		// legal update remote
		{Version{"a": 1, "b": 2}, "a", Version{"a": 1, "b": 3}, true},
		// remote has higher version of self (shouldn't happen in real but let's be sure)
		{Version{"a": 1, "b": 2}, "a", Version{"a": 2, "b": 3}, false},
		// remote update to unknown local
		{Version{}, "a", Version{"b": 3}, true},
		// local version ahead of no Op from remote
		{Version{"a": 1, "b": 2, "c": 4}, "a", Version{"a": 1, "b": 2}, false},
		// remote tries update
		{Version{"a": 1, "b": 2, "c": 4}, "a", Version{"a": 1, "b": 3}, false},
		// remote tries legal value but without knowing of all peer version (version c is unknown to remote)
		{Version{"a": 1, "b": 2, "c": 4}, "a", Version{"a": 1, "b": 5}, false},
		// anti bug test
		{Version{"a": 2}, "a", Version{"b": 3}, false},
		// should be false because it means the same object was created on two peers
		{Version{"a": 0}, "a", Version{"b": 0}, false},
		// another anti bug test
		{Version{"b": 2}, "a", Version{"a": 2, "b": 3}, false},
		{Version{"a": 1, "b": 2}, "a", Version{"b": 3}, false}}
	for _, test := range testValids {
		got := test.local.Valid(test.remote, test.selfid)
		if got != test.want {
			t.Error("Expected", test.want, "got", got, "on", test)
		}
	}
}
