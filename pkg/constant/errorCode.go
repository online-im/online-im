package constant

type CoreErrorCode uint32

const CoreErrorCode_Network = CoreErrorCode(1)
const CoreErrorCode_UserOffLine = CoreErrorCode(2)
const CoreErrorCode_PublishError = CoreErrorCode(3)
const CoreErrorCode_CoreMessageTypeUnsupported = CoreErrorCode(4)

const CoreErrorCode_Unknown = CoreErrorCode(10)
