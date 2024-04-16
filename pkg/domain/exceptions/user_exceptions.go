package exceptions

type UserError string

func (e UserError) Error() string {
	return string(e)
}

func UserNotFound() error {
	return UserError("USER_NOT_FOUND")
}

func UserAlreadyExists() error {
	return UserError("USER_ALREADY_EXISTS")
}

func UserPasswordWrong() error {
	return UserError("USER_PASSWORD_WRONG")
}
