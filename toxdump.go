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
LoadToxDump loads the toxPeerDump file for the local Tinzenite directory.
*/
func LoadToxDump(root string) (*ToxPeerDump, error) {
	data, err := ioutil.ReadFile(root + "/" + TINZENITEDIR + "/" + LOCALDIR + "/" + SELFPEERJSON)
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
Store the toxPeerDump to the directory.
*/
func (t *ToxPeerDump) Store(root string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(root+"/"+TINZENITEDIR+"/"+LOCALDIR+"/"+SELFPEERJSON, data, FILEPERMISSIONMODE)
}
