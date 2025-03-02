INSERT INTO categories (id, name)
VALUES ('task', 'ДЕЛО'),
       ('contact', 'ЗВОНОК/СООБЩЕНИЕ'),
       ('meeting', 'ВСТРЕЧА')
ON CONFLICT DO NOTHING;