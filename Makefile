DB_DSN?="postgres://postgres:password@localhost:5432/users?sslmode=disable"
MIGRATIONS_DIR?="./user-service/migrations"

up:
	docker-compose up -d
down:
	docker-compose down

migrate_up:
	@migrate -verbose -path ${MIGRATIONS_DIR} -database ${DB_DSN} up

migrate_down:
	@migrate -verbose -path ${MIGRATIONS_DIR} -database ${DB_DSN} down

migrate_create:
	@migrate create -dir ${MIGRATIONS_DIR} -ext sql -seq ${name}

migrate_version:
	@migrate -path ${MIGRATIONS_DIR} -database ${DB_DSN} version

#  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto/user.proto