package phonenumber

import (
	"testing"
)

func TestValid(t *testing.T) {

	type Test struct {
		phoneNumber    string
		expectedResult bool
	}

	testCase := []Test{
		{phoneNumber: "0914", expectedResult: false},
		{phoneNumber: "ddddd", expectedResult: false},
		{phoneNumber: "09364524847", expectedResult: true},
	}

	for _, item := range testCase {
		if IsValid(item.phoneNumber) != item.expectedResult {
			t.Error()
		}
	}
}
