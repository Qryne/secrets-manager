package setups

import (
	"context"

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

		row := tx.QueryRow(ctx, `
			INSERT INTO setups (is_setup_complete) VALUES ($1) RETURNING id, is_setup_complete, destroy_at, created_at, updated_at;
		`, false)

		var record db_gen.Setup
		err = row.Scan(&record.ID,
			&record.IsSetupComplete,
			&record.DestroyAt,
			&record.CreatedAt,
			&record.UpdatedAt)
		if err != nil {
			return err
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
