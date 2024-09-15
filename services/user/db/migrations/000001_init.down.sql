-- Удаление пользователя с именем 'admin'
DELETE FROM
  users
WHERE
  username = 'admin';

-- Удаление таблицы пользователей
DROP TABLE IF EXISTS users;

-- Удаление расширения uuid-ossp
DROP EXTENSION IF EXISTS "uuid-ossp";