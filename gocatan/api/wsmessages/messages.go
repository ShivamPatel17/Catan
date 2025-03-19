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
	Board      interface{} `json:"data"`
	PlayerData interface{} `json:"playerData"`
}

type VertexClickedMessage struct {
	EmbeddedBaseMessage
	Data VertexClickedMessageData `json:"data"`
}

type VertexClickedMessageData struct {
	Id string `json:"id"`
}

type BuildSettlementMessage struct {
	EmbeddedBaseMessage
	Data BuildSettlementMessageData `json:"data"`
}

type BuildSettlementMessageData struct {
	PlayerUuid uuid.UUID `json:"playerUuid"`
	VertexUuid uuid.UUID `json:"vertexUuid"`
}
