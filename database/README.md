# Database: ORM -> sqlc

This project's database layer has evolved significantly as I searched for a robust, Go-idiomatic solution:
* **v1 (ORM):** I initially used `gorm`, but ran into several challenges. I learned that the Go community often avoids heavy, "magic" ORMs in favor of more explicit tools.
* **v2 (DSL):** I migrated to `atlas` (using HCL). While powerful, this felt "clunky" and added another language layer on top of SQL.
* **v3 (sqlc):** I landed on **sqlc** as the best solution. It embraces that SQL is the "right tool for the job" by generating fully type-safe Go code directly from my raw SQL queries. This prevents SQL injection, catches errors at compile-time, and is incredibly fast.
