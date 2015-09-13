package shared

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
)

/*
IsTinzenite checks whether a given path is indeed a valid directory
*/
// TODO detect incomplete dir (no connected peers, etc) or write a validate method
func IsTinzenite(dirpath string) bool {
	value, _ := DirectoryExists(dirpath + "/" + TINZENITEDIR)
	return value
}

/*
IsEncrypted checks whether a given path is indeed a valid directory for an
encrypted peer.
*/
// TODO detect incomplete dir (no connected peers, etc) or write a validate method
func IsEncrypted(dirpath string) bool {
	value, _ := DirectoryExists(dirpath + "/" + ORGDIR)
	return value
}

/*
MakeDotTinzenite creates the directory structure for the .tinzenite directory
including the .tinignore file required for it. The given path is the path to the
directory (NOT the .TINZENITEDIR!).
*/
func MakeDotTinzenite(root string) error {
	root = root + "/" + TINZENITEDIR
	// build directory structure
	err := MakeDirectories(root, ORGDIR+"/"+PEERSDIR, TEMPDIR, REMOVEDIR, LOCALDIR, LOCALDIR+"/"+REMOVESTOREDIR, RECEIVINGDIR)
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
	exists, err := FileExists(filePath.FullPath())
	// if file doesn't exist / error happened we're done
	if !exists || err != nil {
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
RemoveDirContents removes all files within the given path, leaving the directory
as is.
*/
func RemoveDirContents(path string) error {
	allStat, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, stat := range allStat {
		err := os.Remove(path + "/" + stat.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

/*
FileExists checks whether a file at that location exists.
*/
func FileExists(path string) (bool, error) {
	stat, err := os.Lstat(path)
	// sadly any error means it doesn't exist, no way to differentiate easily here
	if err != nil {
		return false, nil
	}
	if stat.IsDir() {
		return false, errors.New("path is not a file")
	}
	return true, nil
}

/*
DirectoryExists checks whether a directory at that location exists.
*/
func DirectoryExists(path string) (bool, error) {
	stat, err := os.Lstat(path)
	// sadly any error means it doesn't exist, no way to differentiate easily here
	if err != nil {
		return false, nil
	}
	if !stat.IsDir() {
		return false, errors.New("path is not a directory")
	}
	return true, nil
}

/*
ObjectExists combines FileExists and DirectoryExists.
*/
func ObjectExists(path string) (bool, error) {
	file, errFile := FileExists(path)
	dir, errDir := DirectoryExists(path)
	exists := dir || file
	// if either exists we can ignore all errors as we satisified the query
	if exists {
		return true, nil
	}
	// false case without errors
	if errFile == nil && errDir == nil {
		return false, nil
	}
	// build unified error
	text := "Errors: "
	if errFile != nil {
		text += "FileExists: " + errFile.Error() + " "
	}
	if errDir != nil {
		text += "DirectoryExists: " + errDir.Error()
	}
	// return errors
	return false, errors.New(text)
}

/*
IsDirectoryEmpty checks whether the given directory is empty.
*/
func IsDirectoryEmpty(path string) (bool, error) {
	isDir, err := DirectoryExists(path)
	if err != nil {
		return false, err
	}
	if !isDir {
		return false, errors.New("directory doesn't exist")
	}
	subFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	if len(subFiles) > 0 {
		return false, nil
	}
	return true, nil
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
GetString poses a request to the user and returns his entry as a string.
*/
func GetString(request string) string {
	fmt.Println(request)
	// read input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

/*
GetInt poses a request to the user and returns his entry as an integer.
*/
func GetInt(request string) int {
	for {
		fmt.Println(request)
		// read input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		value, err := strconv.ParseInt(input, 10, 0)
		if err == nil {
			return int(value)
		}
	}
}
