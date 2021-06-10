package message

import "github.com/online-im/online-im/pkg/constant"

type CoreErrorMessage struct {
	Code    constant.CoreErrorCode `json:"code"`
	Message string                 `json:"message"`
}

type CoreMessagePayload struct {
	Data     []byte `json:"data"`
	FromID   string `json:"from_id"`
	TargetID string `json:"target_id"`
}

type CoreMessage struct {
	Type    constant.CoreMessageType `json:"type"`
	Payload CoreMessagePayload       `json:"send"`
	Err     CoreErrorMessage         `json:"err"`
}
