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

func (um *UpdateMessage) String() string {
	data, _ := json.Marshal(um)
	return string(data)
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

func (rm *RequestMessage) String() string {
	data, _ := json.Marshal(rm)
	return string(data)
}
