package shared

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_FileExists(t *testing.T) {
	// make dir for tests so that we can easily clean up afterwards
	root := makeTempDir("", "root")
	defer removeTemp(root)
	// test true
	tempFile := makeTempFile(root, "one")
	exists, err := FileExists(tempFile)
	if exists == false || err != nil {
		t.Error("Expected file to exist, got", exists, "or", err)
	}
	// test false
	err = os.Remove(tempFile)
	if err != nil {
		t.Fatal("Failed test setup", err)
	}
	exists, err = FileExists(tempFile)
	if exists == true || err != nil {
		t.Error("Expected file to NOT exist, got", exists, "or", err)
	}
	// test that dir isn't caught
	tempDir := makeTempDir(root, "two")
	exists, err = FileExists(tempDir)
	if exists == true || err == nil {
		t.Error("Expected directory to not be valid, got", exists, "or", err)
	}
}

func Test_DirectoryExists(t *testing.T) {
	// make dir for tests so that we can easily clean up afterwards
	root := makeTempDir("", "root")
	defer removeTemp(root)
	// test true
	tempDir := makeTempDir(root, "one")
	exists, err := DirectoryExists(tempDir)
	if exists == false || err != nil {
		t.Error("Expected directory to exist, got", exists, "or", err)
	}
	// test false
	err = os.Remove(tempDir)
	if err != nil {
		t.Fatal("Failed test setup", err)
	}
	exists, err = DirectoryExists(tempDir)
	if exists == true || err != nil {
		t.Error("Expected directory to NOT exist, got", exists, "or", err)
	}
	// test that file isn't caught
	tempFile := makeTempFile(root, "file")
	exists, err = DirectoryExists(tempFile)
	if exists == true || err == nil {
		t.Error("Expected file to not be valid, got", exists, "or", err)
	}
}

func Test_ObjectExists(t *testing.T) {
	// make dir for tests so that we can easily clean up afterwards
	root := makeTempDir("", "root")
	defer removeTemp(root)
	// test true
	tempFile := makeTempFile(root, "file")
	tempDir := makeTempDir(root, "dir")
	// we expect true and no error
	exists, err := ObjectExists(tempFile)
	if exists == false || err != nil {
		t.Error("Expected file to exist, got", exists, "or", err)
	}
	exists, err = ObjectExists(tempDir)
	if exists == false || err != nil {
		t.Error("Expected dir to exist, got", exists, "or", err)
	}
	// test false
	err = os.Remove(tempDir)
	if err != nil {
		t.Fatal("Failed test setup", err)
	}
	err = os.Remove(tempFile)
	if err != nil {
		t.Fatal("Failed test setup", err)
	}
	// we expect false and no error
	exists, err = ObjectExists(tempFile)
	if exists == true || err != nil {
		t.Error("Expected file to NOT exist, got", exists, "or", err)
	}
	exists, err = ObjectExists(tempDir)
	if exists == true || err != nil {
		t.Error("Expected dir to NOT exist, got", exists, "or", err)
	}
	// missing: error case
}

func Test_IsDirectoryEmpty(t *testing.T) {
	// make dir for tests so that we can easily clean up afterwards
	root := makeTempDir("", "root")
	defer removeTemp(root)
	// test empty dir
	empty, err := IsDirectoryEmpty(root)
	if empty == false || err != nil {
		t.Error("Expected dir to be empty, got", empty, "or", err)
	}
	// test non empty dir
	tempFile := makeTempFile(root, "file")
	empty, err = IsDirectoryEmpty(root)
	if empty == true || err != nil {
		t.Error("Expected dir to NOT be empty, got", empty, "or", err)
	}
	// test file instead of dir
	empty, err = IsDirectoryEmpty(tempFile)
	if empty == true || err == nil {
		t.Error("Expected error, got", empty, "or", err)
	}
}

func makeTempFile(path, name string) string {
	file, _ := ioutil.TempFile(path, name)
	return file.Name()
}

func makeTempDir(path, name string) string {
	subdir, _ := ioutil.TempDir(path, name)
	return subdir
}

func removeTemp(path string) {
	os.RemoveAll(path)
}
