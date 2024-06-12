docker run --name mydget-db -e POSTGRES_PASSWORD=mypassword -e
 POSTGRES_USER=mydget -e POSTGRES_DB=mydget -p 5432:5432 -d postgres:latest