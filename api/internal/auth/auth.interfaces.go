package auth

// Repository Interfaces
type IAuthRepository interface {
	CreateUserByEmail(email string) (AuthUserModel, error)
}

// Service Interfaces
type IAuthServices interface {
	InitUserSignup(email string) error
}
