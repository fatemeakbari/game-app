package errormessage

type ErrorMessage = string

const (
	InternalError         ErrorMessage = "internal error"
	PhoneNumberDuplicated ErrorMessage = "phone number already exist!"
)
