package usermysql

import (
	"database/sql"
	"errors"
	"fmt"
	"gameapp/model/usermodel"
)

func (db DB) Register(user usermodel.User) (usermodel.User, error) {

	res, err := db.db.Exec(`insert into users(name, phone_number, password, role) values (?,?,?,?)`, user.Name, user.PhoneNumber, user.Password, uint(user.Role))

	if err != nil {
		return usermodel.User{}, err
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (db DB) IsPhoneNumberExist(phoneNumber string) (bool, error) {

	row := db.db.QueryRow(`select id from users where phone_number = ?`, phoneNumber)
	var id uint
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (db DB) FindUserByPhoneNumber(phoneNumber string) (usermodel.User, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	var user usermodel.User
	var createDate []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return usermodel.User{}, errors.New("user not found")
		}
		return usermodel.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil

}

func (db DB) FindUserById(userId uint) (usermodel.User, error) {
	row := db.db.QueryRow(`select * from users where id = ?`, userId)

	var user usermodel.User
	var createDate []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return usermodel.User{}, errors.New("user not found")
		}
		return usermodel.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	return user, nil

}
