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
	r.stack = r.sanitize(root)
	r.limit = len(r.stack)
	return &r
}

/*
CreatePath creates a path directly with a subpath selected. Note that the subpath
is NOT checked if it is relative or absolute: to do that use CreatePathRoot
directly followed by Apply, which will check.
*/
func CreatePath(root string, subpath string) *RelativePath {
	r := CreatePathRoot(root)
	r.stack = append(r.stack, r.sanitize(subpath)...)
	return r
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
	if len(r.stack) == 0 {
		log.Println("RelativePath: WARNING: empty stack!")
		return ""
	}
	// make sure we remain within safe values
	element := len(r.stack) - 1
	if element < 0 {
		element = 0
	}
	return r.stack[element]
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
Apply does two different things depending on the path given. If path begins with
"/" it is considered an absolute path and must match the root of the calling
path. Otherwise the path is applied as a new sub path, replacing the old value.
*/
func (r *RelativePath) Apply(path string) *RelativePath {
	// absolute path?
	if strings.HasPrefix(path, "/") {
		// same root?
		if !strings.HasPrefix(path, r.RootPath()) {
			// if not: return copy of r without applying
			return CreatePath(r.RootPath(), r.SubPath())
		}
		// otherwise set new subpath
		relPath := CreatePathRoot(path)
		relPath.limit = r.limit
		return relPath
	}
	// relative path simply replaces the sub path
	relPath := CreatePath(r.RootPath(), path)
	return relPath
}

/*
RenameLastElement overwrites the last element completely with the given value
and returns the corresponding new RelativePath.
*/
func (r *RelativePath) RenameLastElement(value string) *RelativePath {
	// make sure that we're not renaming the root path
	// also: value must be valid single element, so no '/'!
	if r.Depth() == r.limit || strings.Contains(value, "/") {
		return CreatePath(r.RootPath(), r.SubPath())
	}
	relPath := CreatePath(r.RootPath(), r.SubPath())
	relPath.stack[len(relPath.stack)-1] = value
	return relPath
}

/*
Depth is the amount of elements contained in the full path.
*/
func (r *RelativePath) Depth() int {
	return len(r.stack)
}

/*
AtRoot signals whether the path has any sub path.
*/
func (r *RelativePath) AtRoot() bool {
	return len(r.stack) == r.limit
}

/*
Up removes the last element from the path, up to the root path (and no further!).
*/
func (r *RelativePath) Up() *RelativePath {
	pop := len(r.stack) - 1
	if pop < r.limit {
		pop = r.limit
	}
	return &RelativePath{limit: r.limit, stack: r.stack[:pop]}
}

func (r *RelativePath) String() string {
	return "Path: " + r.FullPath()
}

func (r *RelativePath) sanitize(path string) []string {
	splitted := strings.Split(path, "/")
	var clean []string
	for _, value := range splitted {
		if value != "" {
			clean = append(clean, value)
		}
	}
	return clean
}
