package mysql_test

import (
	"messagingapp/entity"
	"messagingapp/repository/mysql"
	"testing"
)

var db = mysql.New()

func TestRegister(t *testing.T) {

	_, err := db.Register(entity.User{Name: "fateme", PhoneNumber: "09151870482"})

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

		if res, err := db.IsPhoneNumberUnique("0988"); err != nil {
			t.Error(err)
		} else {
			if res != test.expectedResult {
				t.Error()
			}
		}
	}

}
