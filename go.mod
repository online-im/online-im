module github.com/goonline/online-im

go 1.13

require (
	git.go-online.org.cn/Glory/glory v0.0.0-20210605023606-d2b2c86e41db
	github.com/dubbogo/gost v1.11.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	go.uber.org/atomic v1.7.0
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
)

replace git.go-online.org.cn/Glory/glory => ../../../GoOnline/2021/glory
