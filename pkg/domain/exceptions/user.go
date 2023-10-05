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

func UserAlreadyVerified() error {
	return UserError("USER_IS_ALREADY_VERIFIED")
}

func UnverifiedAccount() error {
	return UserError("UNVERIFIED_ACCOUNT")
}

func InvalidCodeResetPassword() error {
	return UserError("INVALID_CODE_TO_RESET_PASSWORD")
}

func InvalidCodeSmsCode() error {
	return UserError("INVALID_SMS_CODE")
}

func InvalidCodeConfirmAccount() error {
	return UserError("INVALID_CODE_TO_CONFIRM_ACCOUNT")
}

func InvalidToken() error {
	return UserError("ACCOUNT_INVALID_TOKEN")
}
