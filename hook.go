package go_sql_hook

import (
	"context"
	"database/sql/driver"
	"fmt"
	"log"
	"strings"
)

func marshalValues(args []driver.Value) []byte {
	res := make([]string, len(args))
	for i, arg := range args {
		res[i] = fmt.Sprintf("%v", arg)
	}
	return []byte(strings.Join(res, ", "))
}

func marshalNameValues(args []driver.NamedValue) []byte {
	res := make([]string, len(args))
	for i, arg := range args {
		res[i] = string(marshalNameValue(&arg))
	}
	return []byte(strings.Join(res, ", "))
}

func marshalNameValue(arg *driver.NamedValue) []byte {
	res := ""
	if arg.Name == "" {
		res = fmt.Sprintf("%s", arg.Value)
	} else {
		res = fmt.Sprintf("%s=%s", arg.Name, arg.Value)
	}
	return []byte(res)
}

type ExecHook func(query string, args []driver.Value) (driver.Result, error)

func defaultExecHook (query string, args []driver.Value) (driver.Result, error) {
	log.Printf("sql-hook: Exec query: %s, args: %s", query, string(marshalValues(args)))
	return driver.ResultNoRows, nil
}

type QueryHook func(query string, args []driver.Value) (driver.Rows, error)

func defaultQueryHook (query string, args []driver.Value) (driver.Rows, error) {
	log.Printf("sql-hook: Query query: %s, args: %s", query, string(marshalValues(args)))
	return &emptyRows{}, nil
}

type PingHook func(ctx context.Context) (err error)

func defaultPingHook(ctx context.Context) (err error) {
	log.Printf("sql-hook: Ping", )
	return nil
}

type QueryContextHook func(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error)

func defaultQueryContextHook(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error){
	log.Printf("sql-hook: QueryContext query: %s, args: %v", query, marshalNameValues(args))
	return &emptyRows{}, nil
}

type ExecContextHook func(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error)

func defaultExecContextHook (ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	log.Printf("sql-hook: ExecContext query: %s, args: %s", query, marshalNameValues(args))
	return driver.ResultNoRows, nil
}

type CheckNamedValueHook func(nv *driver.NamedValue) (err error)

func defaultCheckNamedValueHook (nv *driver.NamedValue) (err error) {
	log.Printf("sql-hook: CheckNamedValue: %s", marshalNameValue(nv))
	return nil
}

type CommitHook func() (err error)

func defaultCommitHook () (err error) {
	log.Printf("sql-hook: Commit")
	return nil
}

type RollbackHook func() (err error)

func defaultRollbackHook () (err error) {
	log.Printf("sql-hook: Rollback")
	return nil
}

type emptyRows struct {

}

func (e emptyRows) Columns() []string {
	return []string{}
}

func (e emptyRows) Close() error {
	return nil
}

func (e emptyRows) Next(dest []driver.Value) error {
	return fmt.Errorf("not data")
}
