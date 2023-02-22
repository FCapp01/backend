// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createMenuStmt, err = db.PrepareContext(ctx, createMenu); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMenu: %w", err)
	}
	if q.createRestaurantStmt, err = db.PrepareContext(ctx, createRestaurant); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRestaurant: %w", err)
	}
	if q.deleteMenuStmt, err = db.PrepareContext(ctx, deleteMenu); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMenu: %w", err)
	}
	if q.deleteRestaurantStmt, err = db.PrepareContext(ctx, deleteRestaurant); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRestaurant: %w", err)
	}
	if q.getMenuStmt, err = db.PrepareContext(ctx, getMenu); err != nil {
		return nil, fmt.Errorf("error preparing query GetMenu: %w", err)
	}
	if q.getRestaurantStmt, err = db.PrepareContext(ctx, getRestaurant); err != nil {
		return nil, fmt.Errorf("error preparing query GetRestaurant: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createMenuStmt != nil {
		if cerr := q.createMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMenuStmt: %w", cerr)
		}
	}
	if q.createRestaurantStmt != nil {
		if cerr := q.createRestaurantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRestaurantStmt: %w", cerr)
		}
	}
	if q.deleteMenuStmt != nil {
		if cerr := q.deleteMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMenuStmt: %w", cerr)
		}
	}
	if q.deleteRestaurantStmt != nil {
		if cerr := q.deleteRestaurantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRestaurantStmt: %w", cerr)
		}
	}
	if q.getMenuStmt != nil {
		if cerr := q.getMenuStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMenuStmt: %w", cerr)
		}
	}
	if q.getRestaurantStmt != nil {
		if cerr := q.getRestaurantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRestaurantStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createMenuStmt       *sql.Stmt
	createRestaurantStmt *sql.Stmt
	deleteMenuStmt       *sql.Stmt
	deleteRestaurantStmt *sql.Stmt
	getMenuStmt          *sql.Stmt
	getRestaurantStmt    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createMenuStmt:       q.createMenuStmt,
		createRestaurantStmt: q.createRestaurantStmt,
		deleteMenuStmt:       q.deleteMenuStmt,
		deleteRestaurantStmt: q.deleteRestaurantStmt,
		getMenuStmt:          q.getMenuStmt,
		getRestaurantStmt:    q.getRestaurantStmt,
	}
}
