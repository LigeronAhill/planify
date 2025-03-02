CREATE UNIQUE INDEX IF NOT EXISTS tasks_statuses_index
    ON tasks_statuses (task_id, status);

