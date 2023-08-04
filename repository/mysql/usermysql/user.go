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
	rows, err := db.db.Query(`select * from users where id = ?`, userId)

	if err != nil {
		return usermodel.User{}, fmt.Errorf("unexpected error : %w", err)
	}

	var user usermodel.User
	for rows.Next() {

		user, err = scanUser(rows)

		if err != nil {
			return usermodel.User{}, fmt.Errorf("unexpected error : %w", err)
		}

		return user, nil

	}
	return usermodel.User{}, errors.New("user not found")

}

func (db DB) UserList() ([]usermodel.User, error) {

	users := make([]usermodel.User, 0)

	rows, err := db.db.Query(`select * from users`)

	//TODO refactor error handling
	if err != nil {
		return users, fmt.Errorf("unexpected error : %w", err)
	}

	var user usermodel.User
	for rows.Next() {

		user, err = scanUser(rows)

		//TODO refactor error handling
		if err != nil {
			return users, fmt.Errorf("unexpected error : %w", err)
		}

		users = append(users, user)

	}
	return users, nil

}

func scanUser(rows *sql.Rows) (usermodel.User, error) {
	var user usermodel.User
	var createDate []uint8
	err := rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createDate, &user.Password, &user.Role)

	return user, err

}
