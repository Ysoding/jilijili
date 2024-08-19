
# ==============================================================================
# Administration

# Define environment variables
JILI_DB_HOST = localhost
JILI_DB_NAME = jili
JILI_DB_PASSWORD = 123456

export JILI_DB_HOST
export JILI_DB_NAME
export JILI_DB_PASSWORD


run:
	go run main.go

migrate:
	go run cmd/tooling/main.go migrate

seed: migrate
	go run cmd/tooling/main.go seed

useradd:
	go run cmd/tooling/main.go useradd x y z

curl-create:
	curl -il -X POST \
	-H 'Content-Type: application/json' \
	-d '{"name":"xmchx","email":"xmchx@test.com","password":"1234#test","passwordConfirm":"1234#test"}' \
	http://localhost:9000/v1/users



ping:
	curl -il http://localhost:9000/v1/ping
