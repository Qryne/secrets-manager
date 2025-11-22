package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PSQLHandler struct {
	Conn *pgx.Conn
}

func (handler *PSQLHandler) Execute(statement string) {
	ctx := context.Background()
	handler.Conn.Exec(ctx, statement)
}

func (handler *PSQLHandler) Query(statement string, args ...any) (IRow, error) {
	ctx := context.Background()

	rows, err := handler.Conn.Query(ctx, statement, args...)

	if err != nil {
		fmt.Println(err)
		return new(PSQLRow), err
	}
	row := new(PSQLRow)
	row.Rows = rows

	return row, nil
}

type PSQLRow struct {
	Rows pgx.Rows
}

func (r PSQLRow) Scan(dest ...any) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r PSQLRow) Next() bool {
	return r.Rows.Next()
}
