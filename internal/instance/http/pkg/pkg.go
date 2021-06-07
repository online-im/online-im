package pkg

type Message struct {
	Message     string `json:"message"`
	PublishType uint32 `json:"publish_type"`
	FromID      string `json:"from_id"`
	TargetID    string `json:"target_id"`
}
