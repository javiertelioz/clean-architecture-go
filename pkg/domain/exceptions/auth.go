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

func InvalidCodeSmsCode() error {
	return AuthError("INVALID_SMS_CODE")
}

func InvalidCodeConfirmAccount() error {
	return AuthError("INVALID_CODE_TO_CONFIRM_ACCOUNT")
}

func InvalidToken() error {
	return AuthError("ACCOUNT_INVALID_TOKEN")
}
