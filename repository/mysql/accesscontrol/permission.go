package accesscontrolmysql

import (
	"database/sql"
	model "gameapp/model/accesscontrol"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	"strings"
)

func (db *DB) GetPermissionByIDs(permissionIDs []uint) ([]model.Permission, error) {

	op := "GetPermissionByIDs"
	permissions := make([]model.Permission, 0)

	if len(permissionIDs) == 0 {
		return permissions, nil
	}
	args := make([]any, len(permissionIDs))
	for i, id := range permissionIDs {
		args[i] = id
	}

	query := `select * from permissions where  id in (?` + strings.Repeat(`,?`, len(permissionIDs)-1) + `)`
	rows, err := db.db.Query(query, args...)

	if err != nil {
		return permissions, perBuildError(op, err)
	}

	for rows.Next() {

		permission, err := scanPermission(rows)
		if err != nil {
			return permissions, perBuildError(op, err)
		}

		permissions = append(permissions, *permission)

	}

	return permissions, nil

}

func scanPermission(rows *sql.Rows) (*model.Permission, error) {
	var permission model.Permission
	err := rows.Scan(&permission.ID, &permission.Title)

	return &permission, err
}

func perBuildError(op string, err error) error {
	return errorhandler.New().
		WithWrappedError(err).
		WithOperation(op).
		WithCodeStatus(errorcodestatus.InternalError).
		WithMessage(errormessage.InternalError)
}
