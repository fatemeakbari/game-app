package phonenumber

import "regexp"

func IsValid(phoneNumber string) bool {

	reg, _ := regexp.Compile(`^09\d{9}`)

	return reg.MatchString(phoneNumber)
}
