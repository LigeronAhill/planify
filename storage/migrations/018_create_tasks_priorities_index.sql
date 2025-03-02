CREATE UNIQUE INDEX IF NOT EXISTS tasks_priorities_index
    ON tasks_priorities (task_id, priority);

