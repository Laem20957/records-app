build:
	docker-compose build records-app

run:
	docker-compose up records-app

migrate:
	migrate -path ./ddl -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
	
swag:
	swag init -g cmd/main.go
