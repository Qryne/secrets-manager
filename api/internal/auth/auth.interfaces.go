package auth

type IAuthRepository interface {
	CreateUserByEmail(email string) (AuthUserModel, error)
}

type IAuthServices interface {
	InitUserSignup(email string) error
}
