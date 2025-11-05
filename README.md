# MusMGR: A Music Event & Repertoire Manager

MusMGR[^1] is an in-progress, full-stack application focusing on musicians and ensembles event and repertoire management. This project is an exploration in building a modern, containerized web application from scratch.

This has been a long-term passion project, paused and iterated on over ~2 years, evolving through major life changes and, more importantly, through my own technical growth. It is not "finished" because its deployment needs (observability, version controlling, updates) outgrew a simple docker compose.

[^1]: This repository was formerly **SongMGR**. While its original specific intent has been lost to time, the project is now being refactored with a clearer, more focused domain. As such, it is a snapshot of an active development and iteration process.

## Project Goals

* Build a performant, lightweight, and modern full-stack application to manage pieces and events related to them.
* Explore a **Go based backend** for API development.
* Experiment modern frontend frameworks.
* Master container-based local development workflows.
* (recently) Implement type-safe database access in Go.

## Tech Stack

* **Frontend:** [SvelteKit](https://kit.svelte.dev/)
* **Backend:** [Go](https://go.dev/) with the [Gin](https://github.com/gin-gonic/gin) framework
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **Database Tooling:** [sqlc](https://sqlc.dev/) for type-safe, compile-time query generation.
* **Development Environment:** [Docker](https://www.docker.com/) & Docker Compose

---

## Status: A Story of Iteration

This repository showcases a journey through different tech choices in pursuit of a better developer experience and a more robust final product.

### 1. Frontend: React -> SvelteKit
The project was initially prototyped in React. I made the deliberate decision to migrate to **SvelteKit**, finding its approach to reactivity and routing logic much simpler and more integrated. This migration was a valuable exercise in comparing modern frontend architecture.

### 2. Backend: Go (Gin)
The backend is written in **Go** using the Gin framework. This choice was made to explore Go's strengths in building simple, high-performance, and strongly-typed APIs.

### 3. Database: ORM -> sqlc
This project's database layer has evolved significantly as I searched for a robust, Go-idiomatic solution:
* **v1 (ORM):** I initially used `gorm`, but ran into several challenges. I learned that the Go community often avoids heavy, "magic" ORMs in favor of more explicit tools.
* **v2 (DSL):** I migrated to `atlas` (using HCL). While powerful, this felt "clunky" and added another language layer on top of SQL.
* **v3 (sqlc):** I landed on **sqlc** as the best solution. It embraces that SQL is the "right tool for the job" by generating fully type-safe Go code directly from my raw SQL queries. This prevents SQL injection, catches errors at compile-time, and is incredibly fast.
### 4. Containerization
The project includes a multi-stage `Dockerfile` for the Go backend and a `docker-compose.yml` file that orchestrates the entire local development stack (frontend, backend, database). I am particularly proud of this setup, as it allows for one-command local testing and demonstrates a practical, end-to-end container workflow.

### Future Plans

This project's development was paused because it outgrew its original, simple "Docker Compose" deployment model. My immediate priority is to finish building my [Kubernetes homelab cluster ("overworld")](https://www.google.com/search?q=https://github.com/YOUR_USERNAME/overworld-repo-link), which is managed via GitOps (FluxCD).

Once that platform is stable, the next phase for MusMGR is to:
1.  Complete the refactor for the "MusMGR" domain.
2.  Build a declarative, K8s-native deployment pipeline for it (e.g., a Helm chart or Kustomization) to be deployed on that cluster.
            
