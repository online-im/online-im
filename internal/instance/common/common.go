package common

import "github.com/goonline/online-im/internal/constant"

type DataFromClient struct {
	FromID      string
	Message     string
	PublishType constant.PublishMessageType
	TargetID    string // if global, this field is empty
}
