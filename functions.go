package shared

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

/*
IsTinzenite checks whether a given path is indeed a valid directory
*/
// TODO detect incomplete dir (no connected peers, etc) or write a validate method
func IsTinzenite(dirpath string) bool {
	_, err := os.Stat(dirpath + "/" + TINZENITEDIR)
	if err == nil {
		return true
	}
	// NOTE: object may exist but we may not have permission to access it: in that case
	//       we consider it unaccessible and thus return false
	return false
}

/*
MakeDotTinzenite creates the directory structure for the .tinzenite directory
including the .tinignore file required for it. The given path is the path to the
directory (NOT the .TINZENITEDIR!).
*/
func MakeDotTinzenite(root string) error {
	root = root + "/" + TINZENITEDIR
	// build directory structure
	err := MakeDirectories(root, ORGDIR+"/"+PEERSDIR, TEMPDIR, REMOVEDIR, LOCALDIR, RECEIVINGDIR)
	if err != nil {
		return err
	}
	// write required .tinignore file
	return ioutil.WriteFile(root+"/"+TINIGNORE, []byte(TINDIRIGNORE), FILEPERMISSIONMODE)
}

/*
RemoveDotTinzenite directory. Specifically leaves all user files but removes all
Tinzenite specific items.
*/
func RemoveDotTinzenite(path string) error {
	if !IsTinzenite(path) {
		return ErrNotTinzenite
	}
	/* TODO remove from directory list*/
	return os.RemoveAll(path + "/" + TINZENITEDIR)
}

/*
WriteDirectoryList adds the given path to the DIRECTORYLIST file. Will try to
avoid writing the same path multiple times.
*/
func WriteDirectoryList(path string) error {
	filePath, err := directoryListPath()
	if err != nil {
		return err
	}
	// make dir in case that it doesn't exist yet (root path here is the directory)
	err = MakeDirectory(filePath.RootPath())
	if err != nil {
		return err
	}
	lines, err := ReadDirectoryList()
	if err != nil {
		return err
	}
	// only add new entry if it doesn't yet exist
	if !Contains(lines, path) {
		lines = append(lines, path)
		newContent := strings.Join(lines, "\n")
		return ioutil.WriteFile(filePath.FullPath(), []byte(newContent), FILEPERMISSIONMODE)
	}
	// if already exists we're done
	return nil
}

/*
ReadDirectoryList reads all registered Tinzenite directories in the system. May
return an empty listing if none found!
*/
func ReadDirectoryList() ([]string, error) {
	filePath, err := directoryListPath()
	if err != nil {
		return nil, err
	}
	// if file doesn't exist we're done
	if !FileExists(filePath.FullPath()) {
		return []string{}, nil
	}
	bytes, err := ioutil.ReadFile(filePath.FullPath())
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "\n"), nil
}

/*
directoryListPath returns the path where the file lies.
*/
func directoryListPath() (*RelativePath, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	return CreatePath(user.HomeDir+"/.config/tinzenite/", DIRECTORYLIST), nil
}

/*
RandomHash genereates one long random hash.
*/
func RandomHash() (string, error) {
	data := make([]byte, RANDOMSEEDLENGTH)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	_, err = hash.Write(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

/*
NewIdentifier creates a new random hash that is intended as identification
strings for all manner of different objects. Length is IDMAXLENGTH.
*/
func NewIdentifier() (string, error) {
	data := make([]byte, RANDOMSEEDLENGTH)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = hash.Write(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil))[:IDMAXLENGTH], nil
}

/*
ContentHash generates the hash of the content of the given file at path.
*/
func ContentHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	buf := make([]byte, CHUNKSIZE)
	// create hash
	for amount := CHUNKSIZE; amount == CHUNKSIZE; {
		amount, _ = file.Read(buf)
		// log.Printf("Read %d bytes", amount)
		hash.Write(buf)
	}
	// return hex representation
	return hex.EncodeToString(hash.Sum(nil)), nil
}

/*
MakeDirectory creates the path.
*/
func MakeDirectory(path string) error {
	err := os.MkdirAll(path, FILEPERMISSIONMODE)
	// TODO this doesn't seem to work... why not?
	if err == os.ErrExist {
		return nil
	}
	// either successful or true error
	return err
}

/*
MakeDirectories creates a number of directories in the given root path. Useful
if a complete directory tree has to be built at once.
*/
func MakeDirectories(root string, subdirs ...string) error {
	for _, path := range subdirs {
		err := MakeDirectory(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
FileExists checks whether a file at that location exists. Currently also usable
for directories.

TODO differentiate between dir and file?
*/
func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

/*
Contains check whether the string slice contains the given string value.
*/
func Contains(s []string, value string) bool {
	for _, entry := range s {
		if entry == value {
			return true
		}
	}
	return false
}

/*
CountFiles within a directory given by path.
*/
func CountFiles(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return -1, err
	}
	return len(files), nil
}
