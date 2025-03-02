CREATE UNIQUE INDEX IF NOT EXISTS tasks_categories_index
    ON tasks_categories (task_id, category);

