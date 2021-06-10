package constant

type CoreMessageType uint32

const CoreMessageType_Message = CoreMessageType(0)
const CoreMessageType_ErrorMessage = CoreMessageType(1)
const CoreMessageType_Close = CoreMessageType(2)
const CoreMessageType_Manage = CoreMessageType(3)
const CoreMessageType_Ping = CoreMessageType(5)
const CoreMessageType_Pong = CoreMessageType(6)
