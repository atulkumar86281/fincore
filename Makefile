postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
newmigrate:
	migrate create -ext sql -dir db/migrate -seq add_user
migrateup: 
	migrate -path  db/migrate -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1: 
	migrate -path  db/migrate -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path  db/migrate -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path  db/migrate -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
dropdb:
	docker exec -it postgres12 dropdb simple_bank
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
mysql:
	docker run --name mysql8 -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=pass \
	-e MYSQL_DATABASE=simple_bank \
	-d mysql:8
mysqlmigrateup:
	migrate -path db/migrate \
	-database "mysql://root:pass@tcp(localhost:3306)/simple_bank" \
	-verbose up
server:
	go run main.go
mock:
	mockgen -package mockDb -destination db/mock/store.go github/atulkumar0001/Bank/db/sqlc Store
newsetup:
	docker exec -it postgres12 dropdb --if-exists simple_bank
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
	migrate -path  db/migrate -database "postgresql://root:pass@localhost:5432/simple_bank?sslmode=disable" -verbose up
	sqlc generate
	mockgen -package mockDb -destination db/mock/store.go github/atulkumar0001/Bank/db/sqlc Store
	go test -v -cover ./...
	go run main.go
sync:
	sqlc generate
	mockgen -package mockDb -destination db/mock/store.go github/atulkumar0001/Bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mysql mysqlmigrateup server mock newsetup migrateup1 migratedown1 sync