package shared

import (
	"encoding/json"
	"io/ioutil"
)

/*
ToxPeerDump stores the self peer information along with the tox binary data
required for it to work.
*/
type ToxPeerDump struct {
	SelfPeer *Peer
	ToxData  []byte
}

/*
LoadToxDumpFrom loads the toxPeerDump file from the given path.
*/
func LoadToxDumpFrom(path string) (*ToxPeerDump, error) {
	path = path + "/" + SELFPEERJSON
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	toxPeerDump := &ToxPeerDump{}
	err = json.Unmarshal(data, toxPeerDump)
	if err != nil {
		return nil, err
	}
	return toxPeerDump, nil
}

/*
StoreTo the toxPeerDump to the given path.
*/
func (t *ToxPeerDump) StoreTo(path string) error {
	// prepare data to write
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	path = path + "/" + SELFPEERJSON
	return ioutil.WriteFile(path, data, FILEPERMISSIONMODE)
}
