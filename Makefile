#!make
include .env.local

# Atlas
schema-inspect:
	atlas schema inspect \
		-u "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

schema-apply:
	atlas schema apply \
		-u "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?search_path=public&sslmode=disable" \
		--to "file://./db.hcl"

# Docker
local-up:
	docker-compose --profile dev \
		-p db_adminer_dev \
		--env-file .env.local \
		-f ./db_adminer.yaml \
		up -d

local-down:
	docker-compose --profile dev \
		-p db_adminer_dev \
		--env-file .env.local \
		-f ./db_adminer.yaml \
		down
