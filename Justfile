set dotenv-load
set shell := ["bash", "-c"]

export POSTGRES_URL := 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}'

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
    pg_dump \
      -s \
      --format=plain \
      --no-owner \
      --no-privileges \
      -h ${POSTGRES_HOST} \
      -U ${POSTGRES_USER} \
      ${POSTGRES_DB} \
      --exclude-table=public.goose_db_version \
      | sed '/^\\restrict/d; /^\\unrestrict/d' \
      > backend/db/schema.sql

create_models:
    @echo "Creating controllers from schema..."
    cd backend/db && sqlc generate

fe_admin_dev:
    @echo "Starting frontend server..."
    cd frontend && pnpm dev:admin

start_containers STAGE:
    @echo "Starting Docker containers..."
    docker compose --profile {{ STAGE }} up

build_containers STAGE:
    @echo "Building Docker containers..."
    docker compose --profile {{ STAGE }} build

stop_containers STAGE:
    @echo "Stopping Docker containers..."
    docker compose --profile {{ STAGE }} down

populate_db:
    @echo "Populating database using fixtures..."
    cd backend && go run cmd/populate/main.go

dev MODE API_URL:
    cd frontend && API_URL={{ API_URL }} pnpm dev:{{ MODE }}
