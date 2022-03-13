package database

type SqlHandler interface {
	Begin() (Tx, error)
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
	QueryRow(string, ...interface{}) Row
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

type Row interface {
	Scan(...interface{}) error
}

type Tx interface {
	Query(string, ...interface{}) (Rows, error)
	Exec(string, ...interface{}) (Result, error)
	Rollback() error
	Commit() error
}
