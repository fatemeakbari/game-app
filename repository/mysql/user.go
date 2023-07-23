package mysql

import (
	"database/sql"
	"errors"
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

	row := db.db.QueryRow(`select id from users where phone_number = ?`, phoneNumber)
	var id uint
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {

			return true, nil
		}

		return false, err
	}

	return false, errors.New("phone number is duplicated")
}
