package go_sql_hook

import (
	"database/sql"
	"database/sql/driver"
	"log"
)

type HookDriver struct {
}

func init(){
	sql.Register("hook", &HookDriver{})
}

type AllHooks struct {
	ConnHooks
	TxHooks
}

var allHooks AllHooks

func SetHooks(_allHooks AllHooks){
	allHooks = _allHooks
}

func getConnHooks() ConnHooks {
	hooks := allHooks.ConnHooks
	if hooks.ExecHook == nil {
		hooks.ExecHook = defaultExecHook
	}

	if hooks.QueryHook == nil {
		hooks.QueryHook = defaultQueryHook
	}

	if hooks.CheckNamedValueHook == nil {
		hooks.CheckNamedValueHook = defaultCheckNamedValueHook
	}

	if hooks.ExecContextHook == nil {
		hooks.ExecContextHook = defaultExecContextHook
	}

	if hooks.QueryContextHook == nil {
		hooks.QueryContextHook = defaultQueryContextHook
	}

	if hooks.PingHook == nil {
		hooks.PingHook = defaultPingHook
	}

	return hooks
}

func getStmtHooks() StmtHooks {
	hooks := allHooks.StmtHooks
	if hooks.QueryHook == nil {
		hooks.QueryHook = defaultQueryHook
	}

	if hooks.ExecHook == nil {
		hooks.ExecHook = defaultExecHook
	}

	return hooks
}

func getTxHooks() TxHooks {
	hooks := allHooks.TxHooks
	if hooks.CommitHook == nil {
		hooks.CommitHook = defaultCommitHook
	}

	if hooks.RollbackHook == nil {
		hooks.RollbackHook = defaultRollbackHook
	}

	return hooks
}

func (m HookDriver) Open(dsn string) (driver.Conn, error) {
	log.Printf("sql-hook: Open dsn: %s", dsn)
	return newConn(getConnHooks(), getTxHooks()), nil
}

func (m HookDriver) OpenConnector(dsn string) (driver.Connector, error) {
	log.Printf("sql-hook: OpenConnector dsn: %s", dsn)
	return newConnector(getConnHooks(), getTxHooks()), nil
}
