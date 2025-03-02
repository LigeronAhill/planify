CREATE TABLE IF NOT EXISTS tasks
(
    id          INTEGER PRIMARY KEY,
    name        TEXT     NOT NULL,
    description TEXT     NOT NULL,
    created_at  DATETIME NOT NULL,
    deadline    DATETIME NOT NULL
);

