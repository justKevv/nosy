package main

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
