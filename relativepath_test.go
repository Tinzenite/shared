package shared

import "testing"

type testPath struct {
	root string
	sub  string
	want string
}

func TestCreateRelativePath(t *testing.T) {
	test := []testPath{
		{"/a", "b", "/a/b"},
		{"/a/b", "c/d", "/a/b/c/d"},
		{"/a", "/b/", "/a/b"},
		{"a", "b", "/a/b"},
		{"/////a/////", "///b///c/d//", "/a/b/c/d"}}
	for _, set := range test {
		path := CreatePath(set.root, set.sub)
		result := path.FullPath()
		if result != set.want {
			t.Error("Expected", set.want, "got", result)
		}
	}
}
