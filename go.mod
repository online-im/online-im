module github.com/goonline/online-im

go 1.13

require (
	git.go-online.org.cn/Glory/glory v0.0.0-20210605023606-d2b2c86e41db
	github.com/dubbogo/gost v1.11.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
)

replace git.go-online.org.cn/Glory/glory => ../../../GoOnline/2021/glory
