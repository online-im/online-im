module github.com/online-im/online-im

go 1.13

require (
	github.com/dubbogo/gost v1.11.3
	github.com/glory-go/glory v0.0.0-20210609140243-fbde74945816
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	go.uber.org/atomic v1.7.0
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/glory-go/glory => ../../glory-go/glory
