package messages

import "github.com/google/uuid"

type BaseMessage interface {
	GetMessageType() string
	GetPlayerUUID() uuid.UUID
}

type EmbeddedBaseMessage struct {
	MessageType string    `json:"messageType"`
	PlayerUuid  uuid.UUID `json:"playerUuid"` // will need some auth at some point
}

func (e *EmbeddedBaseMessage) GetMessageType() string {
	return e.MessageType
}

func (e *EmbeddedBaseMessage) GetPlayerUUID() uuid.UUID {
	return e.PlayerUuid
}

type PlayerConnecting struct {
	EmbeddedBaseMessage
}

type GameStateMessage struct {
	EmbeddedBaseMessage
	Board interface{} `json:"data"`
}

type VertexClickedMessage struct {
	EmbeddedBaseMessage
	Data VertexClickedMessageData `json:"Data"`
}

type VertexClickedMessageData struct {
	Id string `json:"Id"`
}

type BuildSettlementMessage struct {
	EmbeddedBaseMessage
	Data BuildSettlementMessageData `json:"Data"`
}

type BuildSettlementMessageData struct {
	PlayerUuid uuid.UUID `json:"playerUuid"`
	VertexUuid uuid.UUID `json:"vertexUuid"`
}
