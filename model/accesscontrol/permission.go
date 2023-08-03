package accesscontrolmodel

type Permission struct {
	ID    uint
	Title PermissionTitle
}

type PermissionTitle = string

const (
	UserList = PermissionTitle("user_list")
)
