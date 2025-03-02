INSERT INTO users (id, first_name, last_name, username)
VALUES (?1, ?2, ?3, ?4)
ON CONFLICT(id) DO UPDATE SET first_name = ?2, last_name = ?2, username = ?3;
