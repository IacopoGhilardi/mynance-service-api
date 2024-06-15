#docker command
docker run --name mydget-db -e POSTGRES_PASSWORD=mypassword -e POSTGRES_USER=mydget -e POSTGRES_DB=mydget -p 5432:5432 -d postgres:latest


#MIGRATIONS
migrate -path internal/database/migrations -database 'postgres://iacopoghilardi:yourpassword@localhost:5432/mydget?sslmode=disable' up

## Create migration
migrate create -ext sql -dir internal/database/migrations -seq create_users_table