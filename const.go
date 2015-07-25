package shared

import "errors"

/*
Errors of Tinzenite.
*/
var (
	ErrUnsupported      = errors.New("feature currently unsupported")
	ErrIsTinzenite      = errors.New("already a Tinzenite directory")
	ErrNotTinzenite     = errors.New("path is not valid Tinzenite directory")
	ErrNoTinIgnore      = errors.New("no .tinignore file found")
	ErrUntracked        = errors.New("object is not tracked in the model")
	ErrNilInternalState = errors.New("internal state has illegal NIL values")
	ErrConflict         = errors.New("conflict, can not apply")
	ErrIllegalFileState = errors.New("illegal file state detected")
)

/*
Internal errors of Tinzenite.
*/
var (
	errWrongObject = errors.New("wrong ObjectInfo")
)

// constant value here
const (
	/*RANDOMSEEDLENGTH is the amount of bytes used as cryptographic hash seed.*/
	RANDOMSEEDLENGTH = 32
	/*IDMAXLENGTH is the length in chars of new random identification hashes.*/
	IDMAXLENGTH = 16
	/*KEYLENGTH is the length of the encryption key used for challenges and file encryption.*/
	KEYLENGTH = 256
	/*FILEPERMISSIONMODE used for all file operations.*/
	FILEPERMISSIONMODE = 0777
	/*CHUNKSIZE for hashing and encryption.*/
	CHUNKSIZE = 8 * 1024
)

// Path constants here
const (
	TINZENITEDIR  = ".tinzenite"
	TINIGNORE     = ".tinignore"
	DIRECTORYLIST = "directory.list"
	LOCALDIR      = "local"
	TEMPDIR       = "temp"
	RECEIVINGDIR  = "receiving"
	REMOVEDIR     = "removed"
	ORGDIR        = "org"
	PEERSDIR      = "peers"
	ENDING        = ".json"
	AUTHJSON      = "auth" + ENDING
	MODELJSON     = "model" + ENDING
	SELFPEERJSON  = "self" + ENDING
	BOOTJSON      = "boot" + ENDING
)
