CREATE UNIQUE INDEX IF NOT EXISTS tasks_executors_index
ON executors (task_id, user_id);
