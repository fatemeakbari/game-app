package usermodel

type Role uint

const (
	UserRole  = Role(1)
	AdminRole = Role(2)
)

const (
	UserRoleStr  = "user"
	AdminRoleStr = "admin"
)

func (r Role) String() string {

	switch r {
	case 1:
		return UserRoleStr
	case 2:
		return AdminRoleStr
	default: //TODO
		return ""
	}

}
func MapStrToRole(str string) Role {
	switch str {
	case UserRoleStr:
		return UserRole
	case AdminRoleStr:
		return AdminRole
	default: //TODO
		return Role(0)
	}
}
