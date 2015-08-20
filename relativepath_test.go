package shared

import "testing"

type testCreate struct {
	root string
	sub  string
	want string
}

type testApply struct {
	root  string
	sub   string
	apply string
	want  string
}

type testUp struct {
	root   string
	sub    string
	amount int // how often to call up
	want   string
}

func Test_CreatePathRoot(t *testing.T) {
	testCreateRootPath := []testCreate{
		{"///a///b  c/d", "", "/a/b  c/d"},
		{"", "", "/"},
		{"/!\"§$%&/()=}", "", "/!\"§$%&/()=}"},
		{"a/b/", "", "/a/b"}}
	for _, set := range testCreateRootPath {
		path := CreatePathRoot(set.root)
		result := path.FullPath()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}
func Test_CreatePath(t *testing.T) {
	testCreatePath := []testCreate{
		{"/a", "b", "/a/b"},
		{"/a/b", "c/d", "/a/b/c/d"},
		{"/a", "/b/", "/a/b"},
		{"a", "b", "/a/b"},
		{"/////a/////", "///b///c/d//", "/a/b/c/d"},
		{"", "a/b", "/a/b"},
		{"a///b/c/d", "", "/a/b/c/d"}}
	for _, set := range testCreatePath {
		path := CreatePath(set.root, set.sub)
		result := path.FullPath()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}

func TestRelativePath_Apply(t *testing.T) {
	testApply := []testApply{
		{"/a", "b", "/a/b/c", "/a/b/c"},
		{"/a/b", "c/d", "e/f", "/a/b/e/f"},
		{"/", "", "/a////b", "/a/b"},
		{"", "a/b/c/d/", "e", "/e"},
		{"/a", "/b", "/c/d", "/a/b"}} // tests against different root
	for _, set := range testApply {
		path := CreatePath(set.root, set.sub)
		path = path.Apply(set.apply)
		result := path.FullPath()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}

func TestRelativePath_Up(t *testing.T) {
	testUp := []testUp{
		{"root", "sub", 4, "/root"},
		{"/a/b/c/d", "e//f////g/", 2, "/a/b/c/d/e"},
		{"a/b", "c", 1, "/a/b"},
		{"/", "a/b", 200, "/"}} // extreme test
	for _, set := range testUp {
		path := CreatePath(set.root, set.sub)
		for count := 0; count < set.amount; count++ {
			path = path.Up()
		}
		result := path.FullPath()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}

func TestRelativePath_LastElement(t *testing.T) {
	testElement := []testCreate{ // we can reuse testCreate for this test
		{"/a/b", "", "b"},
		{"a//b", "c", "c"},
		{"/a/b/c", "d/e/f", "f"}}
	for _, set := range testElement {
		path := CreatePath(set.root, set.sub)
		result := path.LastElement()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}
