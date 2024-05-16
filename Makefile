postgresdb:
	docker run --name postgresdb -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=p@ssw0rd -d postgres:16.2
createdb:
	 docker exec -it postgresdb createdb --username=root --owner=root orderdb
dropdb: 
	docker exec -it postgresdb dropdb orderdb
migrateup:
	migrate -path "./migrations" -database "postgresql://root:p@ssw0rd@localhost:5431/orderdb?sslmode=disable" -verbose up 
migratedown:
	migrate -path "./migrations" -database "postgresql://root:p@ssw0rd@localhost:5431/orderdb?sslmode=disable" -verbose down 
test:
	go test -v -cover ./...
docs:
	swag init -g './internal/adapters/server/server.go'
	 

.PHONY: postgres createdb dropdb migrateup migratedown test docs