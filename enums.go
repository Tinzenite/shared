package shared

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
	default:
		return "unknown"
	}
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
Request defines what has been requested.
*/
type Request int

const (
	/*ReNone is default empty request.*/
	ReNone Request = iota
	/*ReObject requests an object.*/
	ReObject
	/*ReModel requests the model.*/
	ReModel
	/*RePeer requests the connected peers peer file.*/
	RePeer
)

func (req Request) String() string {
	switch req {
	case ReNone:
		return "None"
	case ReObject:
		return "Object"
	case ReModel:
		return "Model"
	default:
		return "unknown"
	}
}
