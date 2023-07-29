package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"messagingapp/model"
)

func (db *DB) Register(user model.User) (model.User, error) {

	res, err := db.db.Exec(`insert into users(name, phone_number, password) values (?,?,?)`, user.Name, user.PhoneNumber, user.Password)

	if err != nil {
		return model.User{}, fmt.Errorf("cant not save userservice %w", err)
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

func (db *DB) FindUserByPhoneNumber(phoneNumber string) (model.User, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	var user model.User
	var createDate []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("userservice not found")
		}
		return model.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil

}

func (db *DB) FindUserById(userId uint) (model.User, error) {
	row := db.db.QueryRow(`select * from users where id = ?`, userId)

	var user model.User
	var createDate []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("userservice not found")
		}
		return model.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil

}
