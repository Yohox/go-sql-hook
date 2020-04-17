package go_sql_hook

import (
	"context"
	"database/sql/driver"
)

type connector struct {
	ConnHooks
	TxHooks
}


func newConnector(connHooks ConnHooks, txHooks TxHooks) driver.Connector {
	return &connector{
		connHooks,
		txHooks,
	}
}

// Connect implements driver.Connector interface.
// Connect returns a connection to the database.
func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	return newConn(c.ConnHooks, c.TxHooks), nil
}

// Driver implements driver.Connector interface.
// Driver returns &MySQLDriver{}.
func (c *connector) Driver() driver.Driver {
	return &HookDriver{}
}
