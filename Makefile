postgres:
	docker run --name postgre16 -p 5450:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine
createdb:
	docker exec -it postgre16 createdb --username=root --owner=root daithuvien
dropdb:
	docker exec -it postgre16 dropdb daithuvien
.PHONY: postgres createdb dropdb

