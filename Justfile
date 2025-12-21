set dotenv-load
set shell := ["bash", "-c"]

export POSTGRES_URL := 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_NAME}'

install_deps:
    @echo "Installing dependencies..."
    cd backend && go mod download
    cd frontend && pnpm install

create_migration NAME:
    @echo "Creating new migration: {{ NAME }}"
    cd backend && goose -dir db/migrations postgres {{ POSTGRES_URL }} create {{ NAME }} sql

migrate_db:
    @echo "Running database migrations..."
    cd backend && goose -dir db/migrations postgres {{ POSTGRES_URL }} up

dump_db_schema:
    @echo "Dumping database to schema.sql..."
    pg_dump -s -O -h ${POSTGRES_HOST} -U ${POSTGRES_USER} ${POSTGRES_NAME} --exclude-table=public.goose_db_version > backend/db/schema.sql

create_models:
    @echo "Creating controllers from schema..."
    cd backend/db && sqlc generate

start_frontend:
    @echo "Starting frontend server..."
    cd frontend && pnpm dev

start_containers STAGE:
    @echo "Starting Docker containers..."
    COMPOSE_PROFILES={{ STAGE }} docker-compose up -d

build_containers STAGE:
    @echo "Building Docker containers..."
    COMPOSE_PROFILES={{ STAGE }} docker-compose build

stop_containers STAGE:
    @echo "Stopping Docker containers..."
    COMPOSE_PROFILES={{ STAGE }} docker-compose down
