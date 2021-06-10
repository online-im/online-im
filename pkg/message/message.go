package message

import "github.com/online-im/online-im/pkg/constant"

// CoreErrorMessage is sent from instance to client to collect error message
type CoreErrorMessage struct {
	Code    constant.CoreErrorCode `json:"code"`
	Message string                 `json:"message"`
}

// CoreMessagePayload is Message payload when CoreMessage.Type is constant.CoreMessageType_Message
type CoreMessagePayload struct {
	Data        []byte               `json:"data"`
	FromID      string               `json:"from_id"`
	TargetID    string               `json:"target_id"`
	PublishType constant.PublishType `json:"publish_type"`
}

type ManagePayload struct {
	ManageType constant.ManageType `json:"manage_type"`
	GroupID    string              `json:"group_id"`
	Users      []string            `json:"users"`
	User       string              `json:"user"`
}

type CoreMessage struct {
	Type           constant.CoreMessageType `json:"type"`
	MessagePayload CoreMessagePayload       `json:"message_payload"`
	ManagePayload  ManagePayload            `json:"manage_payload"`
	Err            CoreErrorMessage         `json:"err"`
}
