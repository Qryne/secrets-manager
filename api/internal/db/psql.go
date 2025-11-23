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
	psqlRow := new(PSQLRow)
	psqlRow.Rows = rows

	return psqlRow.Rows, nil
}

func (handler *PSQLHandler) BeginTx(ctx context.Context, txOptions any) (ITx, error) {

	tx, err := handler.Conn.Begin(ctx)

	if err != nil {
		fmt.Println(err)
		return new(PSQLTx), err
	}

	psqlTX := new(PSQLTx)
	psqlTX.Tx = tx
	return psqlTX, nil
}

type PSQLTx struct {
	Tx pgx.Tx
}

func (r PSQLTx) Begin(ctx context.Context) (ITx, error) {
	newTx, err := r.Tx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &PSQLTx{Tx: newTx}, nil
}

func (r PSQLTx) Commit(ctx context.Context) error {
	err := r.Tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r PSQLTx) Rollback(ctx context.Context) error {
	err := r.Tx.Rollback(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r PSQLTx) Exec(ctx context.Context, sql string, args ...any) (ICommandTag, error) {
	cmd, err := r.Tx.Exec(ctx, sql, args...)
	if err != nil {
		return cmd, err
	}
	return cmd, nil
}

func (r PSQLTx) Query(ctx context.Context, sql string, args ...any) (IRows, error) {
	rows, err := r.Tx.Query(ctx, sql, args...)
	if err != nil {
		return rows, err
	}
	return rows, nil
}

func (r PSQLTx) QueryRow(ctx context.Context, sql string, args ...any) IRow {
	row := r.Tx.QueryRow(ctx, sql, args...)
	return row
}

type PSQLRow struct {
	Rows pgx.Rows
	Row  pgx.Row
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
