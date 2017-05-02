# gorest
Implement a restful API with postgres database using golang

docker pull postgres:latest

docker run --name pgdb -p 5432:5432 -d -e POSTGRES_PASSWORD=postgres postgres

open pgadmin 4
create new database dvdrental
restore format custom or tar