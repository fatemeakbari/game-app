package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"messagingapp/entity"
)

func (db *DB) Register(user entity.User) (entity.User, error) {

	res, err := db.db.Exec(`insert into users(name, phone_number, password) values (?,?,?)`, user.Name, user.PhoneNumber, user.Password)

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

	return false, nil
}

func (db *DB) FindUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	var user entity.User
	var createDate []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil

}
