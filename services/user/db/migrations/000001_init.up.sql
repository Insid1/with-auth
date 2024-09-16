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
    '$2a$10$Hj.S01IaftDzkhumLRzrweNAss2RbuyXb7MF5uR/YgItI92.dlC5O'
  );