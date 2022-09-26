### how to start

1) specify the database in the config file (PostgresDatabase, PostgresUser, PostgresPassword)
2) migrate migrations or copy paste query commands (migrations)
3) swag init -g cmd/main.go -o api/docs
4) go run cmd/main.go
5) http://localhost:7077/swagger/index.html#/
