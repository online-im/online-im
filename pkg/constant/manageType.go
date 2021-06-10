package constant

type ManageType uint32

const ManageType_CreateGroup = ManageType(0)
const ManageType_AddUserToGroup = ManageType(1)
const ManageType_RemoveUserFromGroup = ManageType(2)
const ManageType_DeleteGroup = ManageType(3)
