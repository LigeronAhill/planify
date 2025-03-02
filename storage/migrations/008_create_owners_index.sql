CREATE UNIQUE INDEX IF NOT EXISTS tasks_owners_index
ON owners (task_id, user_id);