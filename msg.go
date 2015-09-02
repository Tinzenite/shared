package shared

import "encoding/json"

/*
TODO: check if via Unmarshal Method we can output the enums as strings instead of
integer values. Should be written to enums.go, I guess?
*/

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
	data, _ := json.Marshal(um)
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
	Request        Request
	Identification string
}

/*
CreateRequestMessage is a convenience method for building an instance of the message.
*/
func CreateRequestMessage(req Request, identification string) RequestMessage {
	return RequestMessage{
		Type:           MsgRequest,
		Request:        req,
		Identification: identification}
}

/*
JSON representation of this message.
*/
func (rm *RequestMessage) JSON() string {
	data, _ := json.Marshal(rm)
	return string(data)
}

func (rm *RequestMessage) String() string {
	return "RequestMessage{Type:" + rm.Type.String() + ",Request:" +
		rm.Request.String() + ",Identification:" + rm.Identification + "}"
}

/*
NotifyMessage is used to notify another peer of TODO what?
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
	data, _ := json.Marshal(nm)
	return string(data)
}

func (nm *NotifyMessage) String() string {
	// TODO fix
	return nm.JSON()
}
