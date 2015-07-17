package shared

import "testing"

type testCreate struct {
	root string
	sub  string
	want string
}

func TestRelativePathCreate(t *testing.T) {
	t.Log("Testing CreatePathRoot")
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
	t.Log("Testing CreatePath")
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

func TestRelativePathApply(t *testing.T) {
	t.Log("Testing Apply")
}
