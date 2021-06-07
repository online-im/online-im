package message

type Message struct {
	Data     string `json:"data"`
	FromID   string `json:"from_id"`
	TargetID string `json:"target_id"`
	Type     uint32 `json:"type"`
}
