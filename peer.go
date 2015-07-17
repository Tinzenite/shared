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
	Name           string
	Address        string
	Protocol       Communication
	encrypted      bool
	Identification string
	initialized    bool
}

/*
loadPeers loads all peers for the given tinzenite root path.
*/
func loadPeers(root string) ([]*Peer, error) {
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
JSON representation of peer.
*/
func (p *Peer) store(root string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	path := root + "/" + TINZENITEDIR + "/" + ORGDIR + "/" + PEERSDIR + "/" + p.Identification + ENDING
	return ioutil.WriteFile(path, data, FILEPERMISSIONMODE)
}
