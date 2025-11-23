package auth

import (
	"log"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/qryne/api/internal/db"
)

type AuthRepoWithCktBrkr struct {
	AuthRepo AuthRepo
}

func (repo *AuthRepoWithCktBrkr) CreateUserByEmail(email string) (AuthUserModel, error) {
	output := make(chan AuthUserModel, 1)
	hystrix.ConfigureCommand("FindUserByEmail", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("FindUserByEmail", func() error {

		user, _ := repo.AuthRepo.CreateUserByEmail(email)

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return AuthUserModel{}, err
	}
}

type AuthRepo struct {
	DBHandler db.IDbHandler
}

func (repo *AuthRepo) CreateUserByEmail(email string) (AuthUserModel, error) {
	row, err := repo.DBHandler.Query("INSERT INTO users (email, is_email_verified) VALUES($1, $2) RETURNING *", email, true)

	if err != nil {
		log.Fatal(err)
		return AuthUserModel{}, err
	}

	var user AuthUserModel

	row.Scan(&user)

	return user, nil
}
