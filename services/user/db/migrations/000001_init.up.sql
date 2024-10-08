-- Добавление расширения UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Создание таблицы пользователей
CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Добавление пользователя в таблицу
INSERT INTO
  users (username, email, password_hash)
VALUES
  (
    'root',
    'root@root.com',
    -- пароль: 1234 
    '$2a$10$kzelwBVIiF5.FsVjv5.R2.wfUxFyeFzdlk1oiEiOl0TV4trdf0i7C'
  );