package shared

import "testing"

type testPath struct {
	root string
	sub  string
	want string
}

func TestCreateRelativePath(t *testing.T) {
	t.Log("Testing CreatePathRoot")
	testCreateRootPath := []testPath{
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
	testCreatePath := []testPath{
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
