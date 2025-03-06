INSERT INTO users (user_id, first_name, last_name, username, created, updated)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT (user_id) DO UPDATE SET first_name = EXCLUDED.first_name,
last_name = EXCLUDED.last_name,
username = EXCLUDED.username,
updated = EXCLUDED.updated
RETURNING *;
