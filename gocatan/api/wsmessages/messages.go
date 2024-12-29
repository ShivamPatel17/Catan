package types

import "github.com/google/uuid"

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

type BuildSettlementMessage struct {
	BaseMessage
	Data BuildSettlementMessageData `json:"Data"`
}

type BuildSettlementMessageData struct {
	PlayerUuid uuid.UUID `json:"playerUuid"`
	VertexUuid uuid.UUID `json:"vertexUuid"`
}
