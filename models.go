package main

import "github.com/egoist/mygo/ui"

// screen is the screen the window shows: the start screen, where connections
// are chosen, or the workspace of the connected database.
type screen uint8

const (
	screenStart screen = iota
	screenWorkspace
)

// driver is a database engine a connection can use: its name, the scheme of
// its connection strings, its version, the short tag of the version the top
// bar shows, its wire protocol, its default port and its icon. A driver
// without a port, as SQLite, has an empty one.
type driver struct {
	name, scheme, version, tag, wire, port string
	icon                                   *ui.SVG
}

// connection is a saved or recently used connection. status is ok, idle or
// err; latency and ago tell how fast it answered and how long ago it was used.
type connection struct {
	name, driver, host, port, database, user string
	status, latency, ago                     string
}

// table is one table of the database the window shows. The rows are sample
// data until a connection is implemented.
type table struct {
	name    string
	columns []column
	rows    [][]string
}

// column is a column of a table and its SQL type.
type column struct{ name, typ string }

// view is a view, a materialized view or a function of the database, shown
// in the sidebar. kind is its type, as the sidebar shows it.
type view struct{ name, kind string }
