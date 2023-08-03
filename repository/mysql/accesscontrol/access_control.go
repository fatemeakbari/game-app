package accesscontrolmysql

import (
	"database/sql"
	model "gameapp/model/accesscontrol"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
)

func (db *DB) GetUserPermissionTitles(userId, roleId uint) ([]model.PermissionTitle, error) {

	op := "GetUserPermissionTitles"
	permissionIds := make([]uint, 0)

	rows, err := db.db.Query(`select * from access_controls where (actor_id = ? and actor_type = ?) or (actor_id = ? and actor_type = ?)`, userId, model.UserActorType, roleId, model.RoleActorType)

	if err != nil {

		return nil, aclBuildError(op, err)
	}

	for rows.Next() {

		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, aclBuildError(op, err)
		}

		permissionIds = append(permissionIds, acl.PermissionId)
	}

	permissions, err := db.GetPermissionByIDs(permissionIds)

	if err != nil {
		return nil, aclBuildError(op, err)
	}

	permissionTitles := make([]model.PermissionTitle, 0)

	for _, p := range permissions {
		permissionTitles = append(permissionTitles, p.Title)
	}

	return permissionTitles, nil
}

func scanAccessControl(rows *sql.Rows) (*model.AccessControl, error) {
	var acl model.AccessControl
	var createAt []uint8

	err := rows.Scan(&acl.ID, &acl.ActorId, &acl.ActorType, &acl.PermissionId, &createAt)

	return &acl, err
}

func aclBuildError(op string, err error) error {
	return errorhandler.New().
		WithWrappedError(err).
		WithOperation(op).
		WithCodeStatus(errorcodestatus.InternalError).
		WithMessage(errormessage.InternalError)
}
