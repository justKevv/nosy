package main

func sampleTables() []table {
	return []table{
		{
			name:    "users",
			columns: []column{{"id", "bigint"}, {"name", "text"}, {"email", "text"}, {"created_at", "timestamptz"}},
			rows: [][]string{
				{"1", "Ada Lovelace", "ada@example.com", "2026-01-04"},
				{"2", "Alan Turing", "alan@example.com", "2026-01-09"},
				{"3", "Grace Hopper", "grace@example.com", "2026-02-17"},
				{"4", "Linus Torvalds", "linus@example.com", "2026-03-02"},
				{"5", "Margaret Hamilton", "margaret@example.com", "2026-03-21"},
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
