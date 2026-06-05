CREATE TABLE IF NOT EXISTS epics (
                                     id SERIAL PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     description TEXT DEFAULT '',
                                     status TEXT NOT NULL DEFAULT 'Новая'
);

CREATE TABLE IF NOT EXISTS subtasks (
                                        id SERIAL PRIMARY KEY,
                                        title TEXT NOT NULL,
                                        description TEXT DEFAULT '',
                                        status TEXT NOT NULL DEFAULT 'Новая',
                                        epic_id INTEGER NOT NULL REFERENCES epics(id) ON DELETE CASCADE
    );