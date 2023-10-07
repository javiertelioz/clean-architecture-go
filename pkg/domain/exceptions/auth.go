package exceptions

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

func UserAlreadyVerified() error {
	return UserError("USER_IS_ALREADY_VERIFIED")
}

func UnverifiedAccount() error {
	return UserError("UNVERIFIED_ACCOUNT")
}

func InvalidCodeResetPassword() error {
	return AuthError("INVALID_CODE_TO_RESET_PASSWORD")
}
