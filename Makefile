postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb: 
	sudo docker exec -it postgres createdb --username=root --owner=root users

migrateup: 
	migrate -path migration/ -database "postgresql://root:root@localhost:5432/users?sslmode=disable" -verbose up

run:
	docker-compose up

redis:
	docker run -d --name counter -p 6379:6379 redis

test:
	go test -v -cover -short ./...