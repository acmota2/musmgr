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

### Future Plans

This project's development was paused because it outgrew its original, simple "Docker Compose" deployment model. My immediate priority is to finish building my [Kubernetes homelab cluster ("overworld")](https://github.com/acmota2/overworld), which is managed via GitOps (FluxCD).

Once that platform is stable, the next phase for MusMGR is to:
1.  Complete the refactor for the "MusMGR" domain.
2.  Build a declarative, K8s-native deployment pipeline for it (e.g., a Helm chart or Kustomization) to be deployed on that cluster.
            
