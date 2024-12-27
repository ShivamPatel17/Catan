package types

type BaseMessage struct {
	MessageType string `json:"messageType"`
}

type VertexClickedMessage struct {
	BaseMessage
	Data VertexClickedMessageData `json:"Data"`
}

type VertexClickedMessageData struct {
	Id string `json:"Id"`
}

type GameStateMessage struct {
	BaseMessage
	Board interface{} `json:"data"`
}
