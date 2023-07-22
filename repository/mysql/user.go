package mysql

import (
	"database/sql"
	"fmt"
	"messagingapp/entity"
)

func (db *DB) Register(user entity.User) (entity.User, error) {

	res, err := db.db.Exec(`insert into users(name, phone_number) values (?,?)`, user.Name, user.PhoneNumber)

	if err != nil {
		return entity.User{}, fmt.Errorf("cant not save user %w", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (db *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {

			return true, nil
		}

		return false, err
	}

	return false, nil
}
