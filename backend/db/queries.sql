-- name: CreateWork :exec
insert into musmgr.works (id, composed_at, instrumentation, title, create_at, updated_at)
values ($1, $2, $3, $4, now(), now());

-- name: CreateEvent :exec
insert into musmgr.events (id, date, description, event_type, create_at, updated_at)
values ($1, $2, $3, $4, now(), now());

-- name: CreateWorkEvent :exec
insert into musmgr.works_events (work_id, event_id)
values ($1, $2);

-- name: CreateFile :exec
insert into musmgr.files (id, name, work_id, create_at, updated_at)
values ($1, $2, $3, now(), now());

-- name: DeleteWorkEvent :exec
delete from musmgr.works_events
where work_id = $1 and event_id = $2;

-- name: DeleteWork :exec
delete from musmgr.works
where id = $1;

-- name: DeleteEvent :exec
delete from musmgr.events
where id = $1;

-- name: DeleteFile :exec
delete from musmgr.files
where id = $1;

-- name: UpdateWork :exec
update musmgr.works
set composed_at = coalesce(sqlc.narg(composed_at), composed_at),
    instrumentation = coalesce(sqlc.narg(instrumentation), instrumentation),
    title = coalesce(sqlc.narg(title), title),
    updated_at = now()
where id = $1;

-- name: UpdateEvent :exec
update musmgr.events
set date = coalesce(sqlc.narg(date), date),
    description = coalesce(sqlc.narg(description), description),
    event_type = coalesce(sqlc.narg(event_type), event_type),
    updated_at = now()
where id = $1;

-- name: UpdateFile :exec
update musmgr.files
set name = coalesce(sqlc.narg(name), name),
    updated_at = now()
where id = $1;

-- name: GetEvents :many
select * from musmgr.events;

-- name: GetWorks :many
select * from musmgr.works;

-- name: GetInstrumentationWorks :many
select * from musmgr.works
where instrumentation = $1;

-- name: GetEventTypeEvents :many
select * from musmgr.events
where event_type = $1;

-- name: GetWorkFiles :many
select * from musmgr.files
where work_id = $1;

-- name: GetEventWorks :many
select id, composed_at, instrumentation, title from musmgr.works
inner join musmgr.works_events
on musmgr.works.id = musmgr.works_events.work_id
where musmgr.works.id = $1;

-- name: GetWorkEvents :many
select id, date, description, event_type from musmgr.events
inner join musmgr.works_events
on musmgr.events.id = musmgr.works_events.event_id
where musmgr.works_events.work_id = $1;
