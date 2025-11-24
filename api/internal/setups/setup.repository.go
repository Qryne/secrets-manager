package setups

import (
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/qryne/api/internal/db"
	db_gen "github.com/qryne/api/internal/db/sqlc"
)

type SetupRepo struct {
	Db db.IDbHandler
}

func (repo *SetupRepo) CreateEntry() (db_gen.Setup, error) {
	output := make(chan db_gen.Setup, 1)
	hystrix.ConfigureCommand("CreateSetup", hystrix.CommandConfig{Timeout: 1000})

	errorChan := hystrix.Go("CreateSetup", func() error {
		ctx := context.Background()

		tx, err := repo.Db.BeginTx(ctx, nil)

		if err != nil {
			return err
		}

		defer tx.Rollback(ctx)

		row, err := tx.Query(ctx, `
			INSERT INTO setups(is_setup_complete) VALUES(false) RETURNING *;
		`)

		if err != nil {
			return err
		}

		defer row.Close()

		var record db_gen.Setup
		if row.Next() {
			err := row.Scan(&record)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("no row returned")

		}

		if err := tx.Commit(ctx); err != nil {
			return err
		}

		output <- record
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errorChan:
		return db_gen.Setup{}, err
	}

}
