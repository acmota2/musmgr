-- name: CreatePiece :exec
insert into musmgr.pieces (id, composed_at, instrumentation, title, create_at, updated_at)
values ($1, $2, $3, $4, now(), now());

-- name: CreateEvent :exec
insert into musmgr.events (id, happened_at, description, event_type, create_at, updated_at)
values ($1, $2, $3, $4, now(), now());

-- name: CreatePieceEvent :exec
insert into musmgr.pieces_events (piece_id, event_id)
values ($1, $2);

-- name: CreateFile :exec
insert into musmgr.files (id, name, piece_id, create_at, updated_at)
values ($1, $2, $3, now(), now());

-- name: DeletePieceEvent :exec
delete from musmgr.pieces_events
where piece_id = $1 and event_id = $2;

-- name: DeletePiece :exec
delete from musmgr.pieces
where id = $1;

-- name: DeleteEvent :exec
delete from musmgr.events
where id = $1;

-- name: DeleteFile :exec
delete from musmgr.files
where id = $1;

-- name: UpdatePiece :exec
update musmgr.pieces
set composed_at = coalesce(sqlc.narg(composed_at), composed_at),
    instrumentation = coalesce(sqlc.narg(instrumentation), instrumentation),
    title = coalesce(sqlc.narg(title), title),
    updated_at = now()
where id = $1;

-- name: UpdateEvent :exec
update musmgr.events
set happened_at = coalesce(sqlc.narg(happened_at), happened_at),
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

-- name: GetPieces :many
select * from musmgr.pieces;

-- name: GetInstrumentationPieces :many
select * from musmgr.pieces
where instrumentation = $1;

-- name: GetEventTypeEvents :many
select * from musmgr.events
where event_type = $1;

-- name: GetPieceFiles :many
select * from musmgr.files
where piece_id = $1;

-- name: GetEventPieces :many
select id, composed_at, instrumentation, title from musmgr.pieces
inner join musmgr.pieces_events
on musmgr.pieces.id = musmgr.pieces_events.piece_id
where musmgr.pieces.id = $1;

-- name: GetPieceEvents :many
select id, happened_at, description, event_type from musmgr.events
inner join musmgr.pieces_events
on musmgr.events.id = musmgr.pieces_events.event_id
where musmgr.pieces_events.piece_id = $1;
