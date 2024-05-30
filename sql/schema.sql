-- somewhat decoupled from the migrations folder - has to be kept in sync...

CREATE TABLE IF NOT EXISTS race_info (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    location TEXT NOT NULL,
    name TEXT NOT NULL,
    event_date INTEGER NOT NULL,
    run_types TEXT NOT NULL,
    created_at INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS race_results (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    run_type TEXT NOT NULL,
    race_id INTEGER NOT NULL,
    start_time INTEGER NOT NULL,
    end_time INTEGER NOT NULL,
    first_name TEXT,
    last_name TEXT,
    race_number INTEGER,
    FOREIGN KEY(race_id) REFERENCES race_info(id)
);
