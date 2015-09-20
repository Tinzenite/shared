package shared

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*
Peer is the communication representation of a Tinzenite peer.
*/
type Peer struct {
	Name           string        // user defined name for the peer
	Address        string        // tox address of the peer
	Protocol       Communication // for now always Tox
	Trusted        bool          // if trusted peer (meaning it must satisfy a challenge)
	Identification string        // internal ID of peer
}

/*
CreatePeer returns a peer object for the given parameters.
*/
func CreatePeer(name string, address string, trusted bool) (*Peer, error) {
	ident, err := NewIdentifier()
	if err != nil {
		return nil, err
	}
	return &Peer{
		Name:           name,
		Address:        address,
		Protocol:       CmTox,
		Trusted:        trusted,
		Identification: ident}, nil
}

/*
LoadPeers loads all peers for the given tinzenite root path.
*/
func LoadPeers(root string) ([]*Peer, error) {
	path := root + "/" + TINZENITEDIR + "/" + ORGDIR + "/" + PEERSDIR
	peersFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var peers []*Peer
	for _, stat := range peersFiles {
		data, err := ioutil.ReadFile(path + "/" + stat.Name())
		if err != nil {
			log.Println("Error loading peer " + stat.Name() + " from disk!")
			continue
		}
		peer := &Peer{}
		err = json.Unmarshal(data, peer)
		if err != nil {
			log.Println("Error unmarshaling peer " + stat.Name() + " from disk!")
			continue
		}
		peers = append(peers, peer)
	}
	return peers, nil
}

/*
StoreTo the given path a JSON representation of peer.
*/
func (p *Peer) StoreTo(path string) error {
	// prepare data to write
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	// add file name and ending
	path = path + "/" + p.Identification + ENDING
	// write
	return ioutil.WriteFile(path, data, FILEPERMISSIONMODE)
}
