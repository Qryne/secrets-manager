package db

type IDbHandler interface {
	Execute(statement string)
	Query(statement string, args ...any) (IRow, error)
}

type IRow interface {
	Scan(dest ...any) error
	Next() bool
}
