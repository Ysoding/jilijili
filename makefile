
# ==============================================================================
# Administration

# Define environment variables
JILI_DB_HOST = localhost
JILI_DB_NAME = jili
JILI_DB_PASSWORD = 123456

export JILI_DB_HOST
export JILI_DB_NAME
export JILI_DB_PASSWORD

migrate:
	go run cmd/tooling/main.go migrate

seed: migrate
	go run cmd/tooling/main.go seed
