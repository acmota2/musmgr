## 0.1.2 (2026-04-03)

### Feat

- **backend**: create bucket if it doesn't exist

## 0.1.1 (2026-02-28)

### Fix

- **frontend**: ensured frontend api variable is dynamic

## 0.1.0 (2026-02-27)

### BREAKING CHANGE

- Overhauled the entire frontend, it's a completely new application
- will invalidate previous backend completely
- No trace of the old frontend
- MusMGR is assumingly the domain now
- older routing expections won't work with this
- Query logic was mutated, so database probably doesn't behave the same way anymore
- categories no longer exist, files don't have types and event types are an enum

### Feat

- **frontend**: added footer
- **frontend**: finished current admin ui
- **frontend**: all pages are done
- **frontend**: all public routes for the frontend are done
- **frontend**: started implementing pieces page
- **frontend**: added pieces view page
- **frontend**: started making homepage responsive
- **frontend**: made events page
- **frontend**: made homepage with admin and public parity
- **frontend**: added editable pages
- **frontend**: started adding routes
- **frontend**: initial new frontend scaffolding
- **frontend**: restarted frontend
- created a database and files population program for the backend
- **backend**: corrected file handlers
- **backend**: connected pdf creation with file creation handler
- **backend**: added pdf preview generation
- **lefthook**: added lefthook for git hooks
- **biomejs**: added biome to frontend
- **backend**: complete overhaul - finished all routes, config and database schema
- **backend**: added access policies
- **backend**: added minio as a storage option
- **backend**: added local storage module
- **database**: sqlc now uses google/uuid
- **backend**: routing overhaul and new configuration expections
- **controller**: added new database query logic
- devx improvements
- **backend**: during refactor added arg and environment awareness
- **backend**: started refactoring for better handling of different components
- added automations for database management
- **frontend**: readapted the frontend for the new schema
- changed the context from songmgr to musmgr

### Fix

- **backend**: corrected classification
- **lefthook**: corrected glob and root assumptions
- **backend**: minor fixes to endpoints overall
- **backend/controller**: fixed file handling and validated pdf integration
- **backend**: small fixes to db and controllers
- **backend**: corrected composer related endpoints
- **backend**: added get composer route
- **backend**: admin files router was not sending :file_id

### Refactor

- **frontend**: removed wild console.log
- **backend**: simplified db schema to accept strings instead of dates
- **backend**: renamed work -> piece and streamlined config
- **database**: changed work -> piece
- **backend**: changed package name
- made the repository editor agnostic
- **backend**: changed variables from song to work
- cleaned left over file
- updating docs and removing old files
