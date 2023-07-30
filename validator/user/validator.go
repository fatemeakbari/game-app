package user

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

type Validator struct {
	userRepository Repository
}

func New(repository Repository) Validator {
	return Validator{
		userRepository: repository,
	}
}
