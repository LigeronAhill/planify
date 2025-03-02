INSERT INTO statuses (id, name)
VALUES ('new', 'НОВАЯ'),
       ('in_progress', 'В РАБОТЕ'),
       ('pending', 'ОТЛОЖЕНА'),
       ('canceled', 'ОТМЕНЕНА'),
       ('done', 'ВЫПОЛНЕНА')
ON CONFLICT DO NOTHING;