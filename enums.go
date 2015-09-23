package shared

import (
	"encoding/json"
	"errors"
)

/*
Communication is an enumeration for the available communication methods
of Tinzenite peers.
*/
type Communication int

const (
	/*CmNone method.*/
	CmNone Communication = iota
	/*CmTox protocol.*/
	CmTox
)

func (cm Communication) String() string {
	switch cm {
	case CmNone:
		return "None"
	case CmTox:
		return "Tox"
	default:
		return "unknown"
	}
}

/*
MsgType is used to define a message type.
*/
type MsgType int

const (
	/*MsgNone default.*/
	MsgNone MsgType = iota
	/*MsgUpdate is an UpdateMessage.*/
	MsgUpdate
	/*MsgRequest is a RequestMessage.*/
	MsgRequest
	/*MsgNotify is a NotifyMessage.*/
	MsgNotify
	/*MsgLock is a LockMessage.*/
	MsgLock
	/*MsgPush is a PushMessage.*/
	MsgPush
	/*MsgChallenge is a ChallengeMessage.*/
	MsgChallenge
)

func (msg MsgType) String() string {
	switch msg {
	case MsgNone:
		return "none"
	case MsgUpdate:
		return "update"
	case MsgRequest:
		return "request"
	case MsgNotify:
		return "notify"
	case MsgLock:
		return "lock"
	case MsgPush:
		return "push"
	case MsgChallenge:
		return "challenge"
	default:
		return "unknown"
	}
}

/*
MarshalJSON overrides json.Marshal for this type.
*/
func (msg *MsgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(msg.String())
}

/*
UnmarshalJSON overrides json.Unmarshal for this type.
*/
func (msg *MsgType) UnmarshalJSON(data []byte) error {
	value := string(data)
	if len(value) <= 1 {
		return errors.New("impossible MsgType: " + value)
	}
	// split ""
	value = value[1 : len(value)-1]
	switch value {
	case "none":
		*msg = MsgNone
	case "update":
		*msg = MsgUpdate
	case "request":
		*msg = MsgRequest
	case "notify":
		*msg = MsgNotify
	case "lock":
		*msg = MsgLock
	case "push":
		*msg = MsgPush
	case "challenge":
		*msg = MsgChallenge
	default:
		return errors.New("invalid MsgType: " + value)
	}
	return nil
}

/*
Operation is the enumeration for the possible protocol operations.
*/
type Operation int

const (
	/*OpUnknown operation.*/
	OpUnknown = iota
	/*OpCreate operation.*/
	OpCreate
	/*OpModify operation.*/
	OpModify
	/*OpRemove operation.*/
	OpRemove
)

func (op Operation) String() string {
	switch op {
	case OpCreate:
		return "create"
	case OpModify:
		return "modify"
	case OpRemove:
		return "remove"
	default:
		return "unknown"
	}
}

/*
ObjectType defines the type of Request or Push.
*/
type ObjectType int

const (
	/*OtNone is default empty request.*/
	OtNone ObjectType = iota
	/*OtObject requests an object.*/
	OtObject
	/*OtModel requests the model.*/
	OtModel
	/*OtPeer requests the connected peers peer file.*/
	OtPeer
)

func (ot ObjectType) String() string {
	switch ot {
	case OtNone:
		return "none"
	case OtObject:
		return "object"
	case OtModel:
		return "model"
	case OtPeer:
		return "peer"
	default:
		return "unknown"
	}
}

/*
MarshalJSON overrides json.Marshal for this type.
*/
func (ot *ObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ot.String())
}

/*
UnmarshalJSON overrides json.Unmarshal for this type.
*/
func (ot *ObjectType) UnmarshalJSON(data []byte) error {
	value := string(data)
	if len(value) <= 1 {
		return errors.New("impossible ObjectType: " + value)
	}
	// split ""
	value = value[1 : len(value)-1]
	switch value {
	case "none":
		*ot = OtNone
	case "object":
		*ot = OtObject
	case "model":
		*ot = OtModel
	case "peer":
		*ot = OtPeer
	default:
		return errors.New("invalid ObjectType: " + value)
	}
	return nil
}

/*
LockAction defines what action a LockMessage is.
*/
type LockAction int

const (
	/*LoNone is the default empty action.*/
	LoNone LockAction = iota
	/*LoRequest is a request for a lock.*/
	LoRequest
	/*LoRelease is a release request for a lock.*/
	LoRelease
	/*LoAccept is used to notify of a successful previous operation.*/
	LoAccept
)

func (lock LockAction) String() string {
	switch lock {
	case LoNone:
		return "none"
	case LoRequest:
		return "request"
	case LoRelease:
		return "release"
	case LoAccept:
		return "accept"
	default:
		return "unknown"

	}
}

/*
MarshalJSON overrides json.Marshal for this type.
*/
func (lock *LockAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(lock.String())
}

/*
UnmarshalJSON overrides json.Unmarshal for this type.
*/
func (lock *LockAction) UnmarshalJSON(data []byte) error {
	value := string(data)
	if len(value) <= 1 {
		return errors.New("impossible LockAction: " + value)
	}
	// split ""
	value = value[1 : len(value)-1]
	switch value {
	case "none":
		*lock = LoNone
	case "request":
		*lock = LoRequest
	case "release":
		*lock = LoRelease
	case "accept":
		*lock = LoAccept
	default:
		return errors.New("invalid LockAction: " + value)
	}
	return nil
}

/*
NotifyType defines what notify action a NotifyMessage is to be.
*/
type NotifyType int

const (
	/*NoNone is the default empty notify.*/
	NoNone NotifyType = iota
	/*NoRemoved is the removed notification.*/
	NoRemoved
	/*NoMissing is the missing notification.*/
	NoMissing
)

func (n NotifyType) String() string {
	switch n {
	case NoNone:
		return "none"
	case NoRemoved:
		return "removed"
	case NoMissing:
		return "missing"
	default:
		return "unknown"
	}
}

/*
MarshalJSON overrides json.Marshal for this type.
*/
func (n *NotifyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

/*
UnmarshalJSON overrides json.Unmarshal for this type.
*/
func (n *NotifyType) UnmarshalJSON(data []byte) error {
	value := string(data)
	if len(value) <= 1 {
		return errors.New("impossible NotifyType: " + value)
	}
	// split ""
	value = value[1 : len(value)-1]
	switch value {
	case "none":
		*n = NoNone
	case "removed":
		*n = NoRemoved
	case "missing":
		*n = NoMissing
	default:
		return errors.New("invalid NotifyType: " + value)
	}
	return nil
}

/*
Cmd is the enum for which operation the program should execute. Satisfies the
Value interface so that it can be used in flag.
*/
type Cmd int

const (
	/*CmdNone is the default empty command.*/
	CmdNone Cmd = iota
	/*CmdCreate is the create command.*/
	CmdCreate
	/*CmdLoad is the load command.*/
	CmdLoad
	/*CmdBootstrap is the bootstrap command.*/
	CmdBootstrap
)

func (c Cmd) String() string {
	switch c {
	case CmdNone:
		return "none"
	case CmdCreate:
		return "create"
	case CmdLoad:
		return "load"
	case CmdBootstrap:
		return "bootstrap"
	default:
		return "unknown"
	}
}

/*
CmdParse parses a string to cmd. If illegal or can not be matched will simply
return cmdNone.
*/
func CmdParse(value string) Cmd {
	switch value {
	case "create":
		return CmdCreate
	case "load":
		return CmdLoad
	case "bootstrap":
		return CmdBootstrap
	default:
		return CmdNone
	}
}
