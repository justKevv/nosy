package main

func sampleTables() []table {
	return []table{
		{
			name:    "users",
			columns: []column{{"id", "bigint"}, {"name", "text"}, {"email", "text"}, {"phone", "text"}, {"country", "char(2)"}, {"plan", "text"}, {"status", "text"}, {"created_at", "timestamptz"}, {"last_login_at", "timestamptz"}},
			rows: [][]string{
				{"1", "Ada Lovelace", "ada@example.com", "+44 20 7946 0001", "GB", "pro", "active", "2026-01-04", "2026-10-26"},
				{"2", "Alan Turing", "alan@example.com", "+44 20 7946 0002", "GB", "team", "active", "2026-01-09", "2026-10-25"},
				{"3", "Grace Hopper", "grace@example.com", "+1 212 555 0103", "US", "pro", "invited", "2026-02-17", "2026-09-30"},
				{"4", "Linus Torvalds", "linus@example.com", "+1 503 555 0104", "US", "free", "active", "2026-03-02", "2026-10-27"},
				{"5", "Margaret Hamilton", "margaret@example.com", "+1 617 555 0105", "US", "team", "suspended", "2026-03-21", "2026-07-14"},
			},
		},
		{
			name:    "orders",
			columns: []column{{"id", "bigint"}, {"user_id", "bigint"}, {"total", "numeric(10,2)"}, {"status", "text"}},
			rows: [][]string{
				{"1001", "1", "42.50", "shipped"},
				{"1002", "3", "13.00", "pending"},
				{"1003", "2", "199.99", "delivered"},
				{"1004", "5", "7.25", "cancelled"},
			},
		},
		{
			name:    "products",
			columns: []column{{"id", "bigint"}, {"sku", "varchar(32)"}, {"name", "text"}, {"price", "numeric(8,2)"}},
			rows: [][]string{
				{"1", "KB-100", "Mechanical keyboard", "89.00"},
				{"2", "MS-220", "Wireless mouse", "29.90"},
				{"3", "MN-027", "27-inch monitor", "249.00"},
			},
		},
	}
}

func sampleViews() []view {
	return []view{
		{"mv_hourly_p99_latency", "mat"},
		{"v_active_locks", "view"},
		{"fn_prune_stale_wal()", "plpgsql"},
	}
}
