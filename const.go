package shared

import (
	"errors"
	"os"
)

/*
Errors of Tinzenite.
*/
var (
	ErrIllegalParameters = errors.New("illegal parameters given")
	ErrUnsupported       = errors.New("feature currently unsupported")
	ErrIsTinzenite       = errors.New("already a Tinzenite directory")
	ErrNotTinzenite      = errors.New("path is not valid Tinzenite directory")
	ErrNoTinIgnore       = errors.New("no .tinignore file found")
	ErrUntracked         = errors.New("object is not tracked in the model")
	ErrNilInternalState  = errors.New("internal state has illegal NIL values")
	ErrConflict          = errors.New("conflict, can not apply")
	ErrIllegalFileState  = errors.New("illegal file state detected")
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
	/*FILEFLAGCREATEAPPEND is the flag required to create a file or append to it if it already exists.*/
	FILEFLAGCREATEAPPEND = os.O_CREATE | os.O_RDWR | os.O_APPEND
	/*CHUNKSIZE for hashing and encryption.*/
	CHUNKSIZE = 8 * 1024
)

// Path constants here
const (
	TINZENITEDIR   = ".tinzenite"
	TINIGNORE      = ".tinignore"
	DIRECTORYLIST  = "directory.list"
	LOCALDIR       = "local"
	TEMPDIR        = "temp"
	RECEIVINGDIR   = "receiving"
	REMOVEDIR      = "removed"
	REMOVECHECKDIR = "check"
	REMOVEDONEDIR  = "done"
	ORGDIR         = "org"
	PEERSDIR       = "peers"
	ENDING         = ".json"
	AUTHJSON       = "auth" + ENDING
	MODELJSON      = "model" + ENDING
	SELFPEERJSON   = "self" + ENDING
	BOOTJSON       = "boot" + ENDING
)

// .tinignore content for .tinzenite directory
const TINDIRIGNORE = "# DO NOT MODIFY!\n/" + LOCALDIR + "\n/" + TEMPDIR + "\n/" + RECEIVINGDIR

// IDMODEL is the model identification used to differentiate models from files.
const IDMODEL = "MODEL"
