package infra

import (
	"asgo/interfaces/database"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	mysql    string = "mysql"
	postgres string = "postgres"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlTx struct {
	Tx *sql.Tx
}

type SqlResult struct {
	Result sql.Result
}

type SqlRows struct {
	Rows *sql.Rows
}

type SqlRow struct {
	Row *sql.Row
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open(postgres, "root:@tcp(db:3306)")
	if err != nil {
		log.Fatalln(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

//SqlHandler
func (handler *SqlHandler) Begin() (database.Tx, error) {
	tx, err := handler.Conn.Begin()
	if err != nil {
		return new(SqlTx), err
	}
	t := new(SqlTx)
	t.Tx = tx
	return t, err
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return row, nil
}

func (handler *SqlHandler) QueryRow(statement string, args ...interface{}) database.Row {
	row := handler.Conn.QueryRow(statement, args...)
	r := new(SqlRow)
	r.Row = row
	return r
}

// Tx
func (tx SqlTx) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := tx.Tx.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	if err := rows.Err(); err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return row, nil
}

func (tx SqlTx) Exec(statement string, args ...interface{}) (database.Result, error) {
	result, err := tx.Tx.Exec(statement, args...)
	if err != nil {
		return new(SqlResult), err
	}
	res := new(SqlResult)
	res.Result = result
	return res, nil
}

func (tx SqlTx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx SqlTx) Commit() error {
	return tx.Tx.Commit()
}

// Result
func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// Rows
func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

func (r SqlRows) Err() error {
	return r.Rows.Err()
}

// Row
func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan()
}
