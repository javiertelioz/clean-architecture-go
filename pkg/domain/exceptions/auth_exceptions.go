package exceptions

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

func AuthBadCredentials() error {
	return AuthError("USER_BAD_CREDENTIALS")
}

func AuthInvalidToken() error {
	return AuthError("INVALID_TOKEN")
}

func AuthExpiredToken() error {
	return AuthError("EXPIRED_TOKEN")
}
