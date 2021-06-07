package http

type ConnRsp struct {
	Address string `json:"address"`
	Ok      bool   `json:"ok"`
}
