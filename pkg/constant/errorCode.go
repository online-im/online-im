package constant

type CoreErrorCode uint32

const CoreErrorCode_Network = CoreErrorCode(1000)
const CoreErrorCode_DB = CoreErrorCode(1001)

const CoreErrorCode_UserOffLine = CoreErrorCode(2000)

const CoreErrorCode_PublishError = CoreErrorCode(3000)
const CoreErrorCode_ParamError = CoreErrorCode(3001)
const CoreErrorCode_CoreMessageTypeUnsupported = CoreErrorCode(4000)

const CoreErrorCode_Unknown = CoreErrorCode(10000)

func (CoreErrorCode) Error() string {
	// TODO: 定义好通用error类型
	return ""
}
