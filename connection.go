package go_sql_hook

import (
	"context"
	"database/sql/driver"
)

type ConnHooks struct {
	StmtHooks
	PingHook
	QueryContextHook
	ExecContextHook
	CheckNamedValueHook
}

type conn struct {
	ConnHooks
	TxHooks
}

func newConn(connHooks ConnHooks, txHooks TxHooks) driver.Conn {
	return &conn{
		connHooks,
		txHooks,
	}
}


func (c *conn) Begin() (driver.Tx, error) {
	return newTx(c.TxHooks), nil
}

func (c *conn) Close() (err error) {
	return nil
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return newStmt(query, c.StmtHooks), nil
}

func (c *conn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return c.ExecHook(query, args)
}

func (c *conn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return c.QueryHook(query, args)
}
// Ping implements driver.Pinger interface
func (c *conn) Ping(ctx context.Context) (err error) {
	return c.PingHook(ctx)
}

// BeginTx implements driver.ConnBeginTx interface
func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return newTx(c.TxHooks), nil
}

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return c.QueryContextHook(ctx, query, args)
}

func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return c.ExecContextHook(ctx, query, args)
}

func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return newStmt(query, c.StmtHooks), nil
}

func (c *conn) CheckNamedValue(nv *driver.NamedValue) (err error) {
	return c.CheckNamedValueHook(nv)
}

func (c *conn) ResetSession(ctx context.Context) error {
	return nil
}
