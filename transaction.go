package go_sql_hook

import "database/sql/driver"

type TxHooks struct {
	CommitHook
	RollbackHook
}

type tx struct {
	TxHooks
}

func newTx(txHooks TxHooks) driver.Tx {
	return &tx{
		txHooks,
	}
}

func (t *tx) Commit() (err error) {
	return t.CommitHook()
}

func (t *tx) Rollback() (err error) {
	return t.RollbackHook()
}
