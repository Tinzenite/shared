package shared

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Message is a base type for only reading out the operation to define the message
type.
*/
type Message struct {
	Type MsgType
}

/*
UpdateMessage contains the relevant information for notifiying peers of updates.
*/
type UpdateMessage struct {
	Type      MsgType
	Operation Operation
	Object    ObjectInfo
}

/*
CreateUpdateMessage is a convenience method for building an instance of the message.
*/
func CreateUpdateMessage(op Operation, obj ObjectInfo) UpdateMessage {
	return UpdateMessage{
		Type:      MsgUpdate,
		Operation: op,
		Object:    obj}
}

/*
JSON representation of this message.
*/
func (um *UpdateMessage) JSON() string {
	data, err := json.Marshal(um)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (um *UpdateMessage) String() string {
	return "UpdateMessage{Type:" + um.Type.String() + ",Operation:" +
		um.Operation.String() + "," + um.Object.String() + "}"
}

/*
RequestMessage is used to trigger the sending of messages or files from other
peers.
*/
type RequestMessage struct {
	Type           MsgType
	ObjType        ObjectType
	Identification string
}

/*
CreateRequestMessage is a convenience method for building an instance of the message.
*/
func CreateRequestMessage(ot ObjectType, identification string) RequestMessage {
	return RequestMessage{
		Type:           MsgRequest,
		ObjType:        ot,
		Identification: identification}
}

/*
JSON representation of this message.
*/
func (rm *RequestMessage) JSON() string {
	data, err := json.Marshal(rm)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (rm *RequestMessage) String() string {
	return "RequestMessage{Type:" + rm.Type.String() + ",ObjType:" +
		rm.ObjType.String() + ",Identification:" + rm.Identification + "}"
}

/*
NotifyMessage is used to notify another peer of completed removals.
*/
type NotifyMessage struct {
	Type           MsgType
	Operation      Operation
	Identification string
}

/*
CreateNotifyMessage is a convenience method for building an instance of the message.
*/
func CreateNotifyMessage(op Operation, identification string) NotifyMessage {
	return NotifyMessage{
		Type:           MsgNotify,
		Operation:      op,
		Identification: identification}
}

/*
JSON representation of this message.
*/
func (nm *NotifyMessage) JSON() string {
	data, err := json.Marshal(nm)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (nm *NotifyMessage) String() string {
	// TODO fix
	return nm.JSON()
}

/*
LockMessage is the message used to lock an encrypted peer.
*/
type LockMessage struct {
	Type   MsgType
	Action LockAction
}

/*
CreateLockMessage is a convenience method for building an instance of the message.
*/
func CreateLockMessage(action LockAction) LockMessage {
	return LockMessage{
		Type:   MsgLock,
		Action: action}
}

/*
JSON representation of this message.
*/
func (lm *LockMessage) JSON() string {
	data, err := json.Marshal(lm)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (lm *LockMessage) String() string {
	return "LockMessage{Type:" + lm.Type.String() + ",Action:" + lm.Action.String() + "}"
}

/*
PushMessage is the message used to notify an encrypted peer of an incomming file
transfer.
*/
type PushMessage struct {
	Type           MsgType
	Identification string
	ObjType        ObjectType
}

/*
CreatePushMessage is a convenience method for building an instance of the message.
*/
func CreatePushMessage(identification string, ot ObjectType) PushMessage {
	return PushMessage{
		Type:           MsgPush,
		Identification: identification,
		ObjType:        ot}
}

/*
JSON representation of this message.
*/
func (pm *PushMessage) JSON() string {
	data, err := json.Marshal(pm)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (pm *PushMessage) String() string {
	return "PushMessage{Type:" + pm.Type.String() + ",Identification:" + pm.Identification +
		",ObjType:" + pm.ObjType.String() + "}"
}

/*
AuthenticationMessage is the message used to authenticate trusted peers.
*/
type AuthenticationMessage struct {
	Type      MsgType
	Encrypted []byte
	Nonce     *[24]byte
}

/*
CreateAuthenticationMessage is a convenience method for building an instance of the message.
*/
func CreateAuthenticationMessage(encrypted []byte, nonce *[24]byte) AuthenticationMessage {
	return AuthenticationMessage{
		Type:      MsgChallenge,
		Encrypted: encrypted,
		Nonce:     nonce}
}

/*
JSON representation of this message.
*/
func (am *AuthenticationMessage) JSON() string {
	data, err := json.Marshal(am)
	if err != nil {
		log.Println("Msg: JSON error:", err)
	}
	return string(data)
}

func (am *AuthenticationMessage) String() string {
	return "AuthenticationMessage{Type:" + am.Type.String() +
		",Encrypted:" + fmt.Sprintf("%+v", am.Encrypted) +
		",Nonce:" + fmt.Sprintf("%+v", am.Nonce) + "}"
}
