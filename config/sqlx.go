package config

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/constants"
)

type Tx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Connection struct {
	db *sqlx.DB
	DB *sql.DB
}

func NewConnection(db *sqlx.DB) *Connection {
	return &Connection{
		db: db,
		DB: db.DB,
	}
}

type TxFn func(ctx context.Context) error

func (c *Connection) WithTransaction(ctx context.Context, fn TxFn) (err error) {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			log.Error().Msgf("Rollback transaction: %s", err)
			tx.Rollback()
		} else {
			log.Debug().Msgf("Commit transaction")
			tx.Commit()
		}
		log.Debug().Msgf("End transaction")
	}()

	ctx = context.WithValue(ctx, constants.TransactionKey, tx)
	err = fn(ctx)

	return
}

func (c *Connection) TrOrDB(ctx context.Context) Tx {
	tx, ok := ctx.Value(constants.TransactionKey).(Tx)
	if !ok {
		log.Debug().Msgf("No database found in context, default db is used")
		return c.db
	}

	log.Debug().Msgf("Database found as transaction in context")

	return tx
}
