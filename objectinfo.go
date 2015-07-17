package shared

import (
	"encoding/json"
	"os"
)

/*
ObjectInfo represents the in model object fully.
*/
type ObjectInfo struct {
	Directory      bool
	Identification string
	Name           string
	Path           string
	Shadow         bool
	Version        Version
	Content        string        `json:",omitempty"`
	Objects        []*ObjectInfo `json:",omitempty"`
}

/*
CreateObjectInfo is a TEST function for creating an object for the specified
parameters.
*/
func CreateObjectInfo(root string, subpath string, selfid string) (*ObjectInfo, error) {
	path := CreatePath(root, subpath)
	// fetch all values we'll need to store
	id, err := NewIdentifier()
	if err != nil {
		return nil, err
	}
	stat, err := os.Lstat(path.FullPath())
	if err != nil {
		return nil, err
	}
	hash := ""
	if !stat.IsDir() {
		hash, err = ContentHash(path.FullPath())
		if err != nil {
			return nil, err
		}
	}
	return &ObjectInfo{
		Directory:      stat.IsDir(),
		Identification: id,
		Name:           path.LastElement(),
		Path:           path.SubPath(),
		Shadow:         false,
		Version:        Version{},
		Content:        hash}, nil
}

/*
Equal checks wether the given pointer points to the same object based on pointer
and identification. NOTE: Does not compare any other properties!
*/
func (o *ObjectInfo) Equal(that *ObjectInfo) bool {
	return o == that || o.Identification == that.Identification
}

/*
String returns a json representation of this object.
*/
func (o *ObjectInfo) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
