CREATE TABLE IF NOT EXISTS tasks (
    task_id INTEGER PRIMARY KEY,
    author_id INTEGER NOT NULL REFERENCES users (user_id),
    executor_id INTEGER NOT NULL REFERENCES users (user_id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    category TEXT NOT NULL,
    created DATETIME NOT NULL,
    updated DATETIME NOT NULL
);
