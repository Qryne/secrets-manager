package apikeys

import (
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/qryne/api/internal/db"
	db_gen "github.com/qryne/api/internal/db/sqlc"
)

type APIKeyRepo struct {
	Db db.IDbHandler
}

func (repo *APIKeyRepo) CreateAPIKey(
	name, slug, prefix, public_id, encryption_iv, encrypted_text, algorithm, setup_id string, scope []string,
) (db_gen.ApiKey, error) {

	output := make(chan db_gen.ApiKey, 1)
	hystrix.ConfigureCommand("CreateAPIKey", hystrix.CommandConfig{Timeout: 1000})

	errChan := hystrix.Go("CreateAPIKey", func() error {
		ctx := context.Background()

		tx, err := repo.Db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx) // safe rollback

		row, err := tx.Query(ctx, `
            INSERT INTO api_keys (
                name, slug, prefix, public_id, encryption_iv,
                encrypted_text, algorithm, setup_id, scope
            )
            VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
            RETURNING id, name, slug, prefix, public_id, encryption_iv,
                      encrypted_text, algorithm, setup_id, scope, created_at
        `,
			name, slug, prefix, public_id, encryption_iv,
			encrypted_text, algorithm, setup_id, scope,
		)
		if err != nil {
			return err
		}
		defer row.Close()

		var record db_gen.ApiKey
		if row.Next() {
			err = row.Scan(
				&record.ID,
				&record.Name,
				&record.Slug,
				&record.Prefix,
				&record.PublicID,
				&record.EncryptionIv,
				&record.EncryptedText,
				&record.Algorithm,
				&record.SetupID,
				&record.Scope,
				&record.CreatedAt,
			)
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
	case err := <-errChan:
		return db_gen.ApiKey{}, err
	}
}
