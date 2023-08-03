package errormessage

type ErrorMessage = string

const (
	InternalError         ErrorMessage = "system internal error"
	PhoneNumberDuplicated ErrorMessage = "phone number already exist"
	LoginMessage          ErrorMessage = "phone number or password is wrong"
	ForbiddenMessage      ErrorMessage = "user access is forbidden"
)
