# 🏥 VitaCare — Sistema de Gestão em Saúde

Servidor HTTP em **GoLang** com CRUD completo de usuários, usando **PostgreSQL** como banco de dados e interface web com tema de saúde.

---

## Estrutura do Projeto

```
servidorHTTP/
├── .env                        # Variáveis de ambiente
├── docker-compose.yml          # PostgreSQL via Docker
├── go.mod                      # Dependências Go
├── schema.sql                  # SQL para criar a tabela
├── app/
│   ├── main.go                 # Ponto de entrada do servidor
│   ├── handlers/
│   │   ├── formHandler.go          # POST /api/create
│   │   ├── loginHandler.go         # POST /api/login
│   │   ├── updateAccountHandler.go # POST /api/update
│   │   └── deleteAccountHandler.go # POST /api/delete
│   └── utils/
│       ├── connectToDB.go      # Conexão com PostgreSQL
│       ├── encrypt.go          # Hash bcrypt de senhas
│       ├── validateUser.go     # Validação de campos
│       ├── getUserByEmail.go   # Busca de usuário
│       ├── updateUser.go       # Atualização no banco
│       └── deleteUser.go       # Remoção do banco
└── static/
    ├── index.html
    ├── forms/
    │   ├── createAccount.html
    │   ├── login.html
    │   ├── updateAccount.html
    │   └── deleteAccount.html
    └── styles/
        ├── index.style.css
        ├── createAccount.style.css
        ├── login.style.css
        ├── updateAccount.style.css
        └── deleteAccount.style.css
```

---

## Configuração

### 1. Pré-requisitos
- [Go 1.21+](https://go.dev/dl/)
- [PostgreSQL 15+](https://www.postgresql.org/) ou [Docker](https://www.docker.com/)

### 2. Clone e configure o `.env`
```bash
git clone <URL_DO_REPOSITORIO>
cd servidorHTTP
cp .env .env.local  # edite com seus dados
```

Conteúdo do `.env`:
```
DB_USER=saude_user
DB_PASSWORD=saude_pass
DB_NAME=saude_db
DB_HOST=localhost
DB_PORT=5432
```

### 3. Suba o banco com Docker
```bash
docker compose up -d
```

### 4. Crie a tabela
```bash
psql -h localhost -U saude_user -d saude_db -f schema.sql
```

### 5. Instale dependências
```bash
go mod tidy
```

### 6. Inicie o servidor
```bash
go run app/main.go
```

Acesse: **http://127.0.0.1:3000**

---

## Rotas

| Método | Rota           | Descrição                    |
|--------|----------------|------------------------------|
| GET    | `/`            | Página inicial               |
| POST   | `/api/create`  | Criar conta                  |
| POST   | `/api/login`   | Login e exibir perfil        |
| POST   | `/api/update`  | Atualizar dados do usuário   |
| POST   | `/api/delete`  | Excluir conta                |

---

## Segurança
- Senhas criptografadas com **bcrypt** (custo padrão)
- Validação de credenciais antes de qualquer operação de update/delete
- Campos validados no backend antes de consultar o banco

---

## Dependências
- [`github.com/lib/pq`](https://github.com/lib/pq) — driver PostgreSQL
- [`github.com/joho/godotenv`](https://github.com/joho/godotenv) — variáveis de ambiente
- [`golang.org/x/crypto`](https://pkg.go.dev/golang.org/x/crypto) — bcrypt
