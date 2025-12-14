# MusMGR: a repertoire manager and tracker

***This repository was formerly **SongMGR**. While its original specific intent has been lost to time, the project is now being refactored and reimagined into a similar, but different domain. It currently finds itself under active (re)development.***

MusMGR is an in-progress, full-stack application focusing on a composer's (or any musician) event and repertoire management. This project is an exploration in building a modern, containerized web application from scratch.

This has been a long-term passion project, paused and iterated on over ~2 years, evolving through major life changes and through my own technical growth. It is not ready yet due to the unseen changes in its original scope, but also because its deployment needs have shifted.

## Project goals

* Build a performant and lightweight full-stack application to manage pieces and events related to them.
* Explore Go as a backend language.
* Originally to explore React, but the current project presents itself with a SvelteKit frontend.
* Master container-based local development workflows.
* (recently) Implement type-safe database access in Go.

## Tech stack

* **Frontend:** [SvelteKit](https://kit.svelte.dev/)
* **Backend:** [Go](https://go.dev/) with the [Gin](https://github.com/gin-gonic/gin) framework
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **Database management:** [sqlc](https://sqlc.dev/) for type-safe, compile-time query generation.
* **Database migrations:** [goose](https://pressly.github.io/goose)
* **Development Environment:** [Docker](https://www.docker.com/) (with Docker Compose)

---

### Future plans

* Complete the refactor from SongMGR to MusMGR.
* 
