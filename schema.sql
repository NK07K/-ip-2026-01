-- Schema do banco de dados VitaCare
-- Execute este arquivo no PostgreSQL para criar a tabela de usuários

CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(100) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    born_date  DATE NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índice para busca por e-mail (otimização)
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
