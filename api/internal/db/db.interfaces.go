package db

import (
	"context"
)

type IDbHandler interface {
	Execute(statement string)
	Query(statement string, args ...any) (IRow, error)
	BeginTx(context context.Context, txOptions any) (ITx, error)
}

type ITx interface {
	Begin(ctx context.Context) (ITx, error)

	Commit(ctx context.Context) error

	Rollback(ctx context.Context) error

	Exec(ctx context.Context, statement string, args ...any) (commandTag ICommandTag, err error)
	Query(ctx context.Context, statement string, args ...any) (IRows, error)
	QueryRow(ctx context.Context, statement string, args ...any) IRow
}

type ICommandTag interface {
	Delete() bool
	Insert() bool
	RowsAffected() int64
	Select() bool
	String() string
	Update() bool
}

type IRows interface {
	Scan(dest ...any) error
	Next() bool
	Close()
}

type IRow interface {
	Scan(dest ...any) error
}
