#DOCKER
docker compose up --build


#MIGRATIONS
migrate -path internal/database/migrations -database 'postgres://iacopoghilardi:yourpassword@localhost:5432/mydget?sslmode=disable' up

## Create migration
migrate create -ext sql -dir internal/database/migrations -seq create_users_table