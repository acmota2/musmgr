-- name: GetWorks :many
select * from works;

-- name: CreateWork :exec
insert into works (id, composed_at, instrumentation, title)
values ($1, $2, $3, $4);

-- name: CreateEvent :exec
insert into events (id, date, description, event_type)
values ($1, $2, $3, $4);

-- name: GetSubcategoryWorks :many
select * from works
where instrumentation = $1;

-- name: GetEventTypeEvents :many
select * from events
where event_type = $1;

-- name: GetWorkFiles :many
select * from files
where work_id = $1;

-- name: GetEventWorks :many
select id, composed_at, instrumentation, title from works
inner join works_events
on works.id = works_events.work_id
where works.id = $1;

-- name: CreateWorkEvent :exec
insert into works_events (work_id, event_id)
values ($1, $2);
