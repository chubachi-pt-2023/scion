$ brew services start postgresql

$ psql postgres

```
postgres=# CREATE DATABASE atomono;
CREATE USER jabelic WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE atomono TO jabelic;
\q
```

## dev

$ go run ./cmd/server/main.go
