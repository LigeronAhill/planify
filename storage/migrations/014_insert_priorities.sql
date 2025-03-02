INSERT INTO priorities (id, name)
VALUES ('high', 'ВЫСОКИЙ'),
       ('normal', 'СРЕДНИЙ'),
       ('low', 'НИЗКИЙ'),
       ('top', 'НАИВЫСШИЙ')
ON CONFLICT DO NOTHING;