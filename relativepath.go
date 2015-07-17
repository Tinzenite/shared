package shared

import (
	"log"
	"strings"
)

/*
RelativePath implements a path consisting of a base path and any subpath that
lies beneath it.
*/
type RelativePath struct {
	stack []string
	limit int
}

/*
CreatePathRoot creates a RelativePath with the given root path.
*/
func CreatePathRoot(root string) *RelativePath {
	r := RelativePath{}
	list := strings.Split(root, "/")
	for _, element := range list {
		if element != "" {
			r.stack = append(r.stack, element)
		}
	}
	r.limit = len(r.stack)
	return &r
}

/*
CreatePath creates a path directly with a subpath selected.
*/
func CreatePath(root string, subpath string) *RelativePath {
	r := CreatePathRoot(root)
	return r.Apply(root + "/" + subpath)
}

/*
FullPath returns the full path of the path.
*/
func (r *RelativePath) FullPath() string {
	return "/" + strings.Join(r.stack, "/")
}

/*
LastElement returns the last element of the complete path.
*/
func (r *RelativePath) LastElement() string {
	return r.stack[len(r.stack)-1]
}

/*
SubPath returns the current sub path.
*/
func (r *RelativePath) SubPath() string {
	return strings.Join(r.stack[r.limit:], "/")
}

/*
RootPath returns the root path.
*/
func (r *RelativePath) RootPath() string {
	return "/" + strings.Join(r.stack[:r.limit], "/")
}

/*
Apply tries to apply the given path to the RelativePath. If it fails it will
return the unmodified RelativePath.

TODO this is wrong. Test it!
*/
func (r *RelativePath) Apply(path string) *RelativePath {
	log.Println("RelativePath.Apply is not yet correct!")
	if strings.HasPrefix(path, r.FullPath()) {
		relPath := CreatePathRoot(path)
		relPath.limit = r.limit
		return relPath
	}
	return &RelativePath{limit: r.limit, stack: r.stack}
}

/*
Depth is the amount of elements contained in the full path.
*/
func (r *RelativePath) Depth() int {
	return len(r.stack)
}

/*
Up removes the last element from the path, up to the root path.

TODO this may be wrong. Test it!
*/
func (r *RelativePath) Up() *RelativePath {
	log.Println("RelativePath.Up may be wrong yet!")
	pop := len(r.stack) - 1
	if pop < r.limit {
		pop = r.limit
	}
	return &RelativePath{limit: r.limit, stack: r.stack[:pop]}
}
