package exceptions

var errors = map[string]string{
	DUPLICATE_EMAIL: "user with the email already exists",
	INVALID_ID:      "couldn't find user with the specified ID",
	NO_USER_FOUND:   "no user found",
}

func ErrorMessage(errorCode string) string {
	return errors[errorCode]
}
