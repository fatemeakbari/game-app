package accesscontrolmodel

type AccessControl struct {
	ID           uint
	ActorId      uint
	ActorType    ActorType
	PermissionId uint
}

type ActorType = string

const (
	UserActorType = ActorType("user")
	RoleActorType = ActorType("role")
)
