package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type dbKey struct{}

func NewContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func FromContext(ctx context.Context) (*sql.DB, bool) {
	value, ok := ctx.Value(dbKey{}).(*sql.DB)
	return value, ok
}

func BeginTx(ctx context.Context) (context.Context, error) {
	db, ok := FromContext(ctx)
	if !ok {
		return ctx, errors.New("not found db from context")
	}
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return ctx, fmt.Errorf("failed to begin tx, %w", err)
	}

}
