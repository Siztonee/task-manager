## How to Run

### 1. Requirements
- Go 1.25.5+
- Postgresql

### 2. Clone
```bash
git clone https://github.com/Siztonee/task-manager.git
cd task-manager
```

### 3. Instal depends
```bash
go mod tidy
```

### 4. Run postgresql and Change the login information for postgre in .env file
```env
  POSTGRES_HOST=localhost
  POSTGRES_PORT=db_port
  POSTGRES_DB=db_name
  POSTGRES_USER=db_user
  POSTGRES_PASSWORD=db_password
  
  DB_SSLMODE=disable
  
  JWT_SECRET=your-super-key
```

### 5. Run server
```bash
go run cmd/api/main.go
```
