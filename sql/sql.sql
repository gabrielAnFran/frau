CREATE TABLE expenses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    expense_name TEXT,
    amount REAL,
    frequency TEXT,
    start_date TEXT,
    end_date TEXT,
    client_id INTEGER,
    data_exc TEXT
);