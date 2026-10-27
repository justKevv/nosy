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

func sampleDrivers() []driver {
	return []driver{
		{name: "PostgreSQL", scheme: "postgres", version: "PostgreSQL 16.4", wire: "wire protocol v3", port: "5432", icon: databaseIcon},
		{name: "MySQL", scheme: "mysql", version: "MySQL 8.4", wire: "wire protocol 10", port: "3306", icon: databaseIcon},
		{name: "SQLite", scheme: "sqlite", version: "SQLite 3.46", wire: "file database", icon: fileIcon},
		{name: "MongoDB", scheme: "mongodb", version: "MongoDB 7.0", wire: "wire protocol 2013", port: "27017", icon: leafIcon},
		{name: "Redis", scheme: "redis", version: "Redis 7.2", wire: "RESP3", port: "6379", icon: boltIcon},
		{name: "SQL Server", scheme: "sqlserver", version: "SQL Server 2022", wire: "TDS 7.4", port: "1433", icon: serverIcon},
	}
}

func sampleConnections() []connection {
	return []connection{
		{name: "prod-us-east-1", driver: "PostgreSQL", host: "prod-us-east-1.db.internal", port: "5432", database: "telemetry_core", user: "ops_readonly", status: "ok", latency: "1.8ms", ago: "2m"},
		{name: "prod-eu-west-1", driver: "MySQL", host: "prod-eu-west-1.db.internal", port: "3306", database: "events_warehouse", user: "ops_readonly", status: "ok", latency: "4.2ms", ago: "1h"},
		{name: "staging-use1", driver: "MongoDB", host: "staging-use1.db.internal", port: "27017", database: "sessions", user: "app_rw", status: "ok", latency: "12ms", ago: "3h"},
		{name: "dev-docker", driver: "SQLite", database: "app.db", status: "idle", latency: "0.4ms", ago: "yest."},
		{name: "cache-primary", driver: "Redis", host: "cache-primary.internal", port: "6379", database: "session_store", status: "err", ago: "2d"},
	}
}

func sampleRecent() []connection {
	return []connection{
		{name: "audit-archive", driver: "PostgreSQL", host: "audit-archive.db.internal", port: "5432", database: "audit_log", user: "ops_readonly", status: "ok", latency: "84ms", ago: "4d"},
		{name: "analytics-replica", driver: "PostgreSQL", host: "analytics-replica.db.internal", port: "5432", database: "analytics", user: "reporting", status: "idle", ago: "1w"},
		{name: "sandbox", driver: "SQLite", database: "sandbox.db", status: "idle", ago: "2w"},
	}
}
