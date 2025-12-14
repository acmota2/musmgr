-- name: GetWorks :many
select * from musmgr.works;

-- name: CreateWork :exec
insert into musmgr.works (id, composed_at, instrumentation, title)
values ($1, $2, $3, $4);

-- name: CreateEvent :exec
insert into musmgr.events (id, date, description, event_type)
values ($1, $2, $3, $4);

-- name: GetSubcategoryWorks :many
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

-- name: CreateWorkEvent :exec
insert into musmgr.works_events (work_id, event_id)
values ($1, $2);
