# MusMGR: a repertoire manager and performance event tracker

MusMGR is a webapp built with a SvelteKit frontend, a Go backend (Gin), and PostgreSQL for persistence. SQL queries are handled via sqlc with goose for migrations, and local development is containerized with Docker Compose.

MusMGR focuses on tracking repertoire and performance events. It is not a music streaming or playback platform.

## Tech stack

* [SvelteKit](https://kit.svelte.dev/) with [Tailwind CSS](https://tailwindcss.com/) (frontend)
* [Go](https://go.dev/) with [Gin](https://github.com/gin-gonic/gin) (backend)
* [PostgreSQL](https://www.postgresql.org/) (database)
* [sqlc](https://sqlc.dev/) (type-safe query generation)
* [goose](https://pressly.github.io/goose)  (database migrations)
* [Docker Compose](https://docs.docker.com/compose) (local development environment)

## Project status

MusMGR is under active development. APIs, data models, and deployment details may change as the project evolves.

---

## Future plans

* Add both MinIO and local storage support for files.
* Implement public and private routes.

---

## Contributors

- **acmota2** - original author and maintainer

Additional contributions may be listed here in the future.
