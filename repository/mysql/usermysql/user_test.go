package usermysql

import (
	"gameapp/model/usermodel"
	"gameapp/repository/mysql"
	"testing"
)

var db = New(mysql.Config{})

func TestRegister(t *testing.T) {

	_, err := db.Register(usermodel.User{Name: "fateme", PhoneNumber: "09151870482"})

	if err != nil {
		t.Error(err)
	}
}

func TestIsPhoneNumberUnique(t *testing.T) {

	type Test struct {
		phoneNumber    string
		expectedResult bool
	}

	testcases := []Test{
		{phoneNumber: "09151870482", expectedResult: false},
		{phoneNumber: "09151444444", expectedResult: true},
	}

	for _, test := range testcases {

		if res, err := db.IsPhoneNumberExist("0988"); err != nil {
			t.Error(err)
		} else {
			if res != test.expectedResult {
				t.Error()
			}
		}
	}

}
