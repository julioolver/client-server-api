CREATE TABLE quotations (
    id TEXT PRIMARY KEY,
    bid REAL NOT NULL,
    name TEXT,
    created_at DATETIME DEFAULT (DATETIME('now'))
);