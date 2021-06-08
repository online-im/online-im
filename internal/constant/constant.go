package constant

type PublishMessageType uint32

const Publish2User = PublishMessageType(1)
const Publish2Room = PublishMessageType(2)
const Publish2Global = PublishMessageType(3)

const GRPCProviderName = "OnlineIMGRPCProviderName"
const IMInstanceServiceID = "im-instance"

const GRPCClientName = "OnlineIMGRPCClientName"

const HTTPProviderName = "OnlineIMHTTPProviderName"

const RegistryKey = "registryKey"
const RedisKey = "redisKey"

const IMInstanceIPsKey = "instance-ips"

// gateway
const HTTPGatewayProviderName = "OnlineIMGatewayHTTPProviderName"
