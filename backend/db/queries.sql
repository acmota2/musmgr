-- name: CreateComposer :exec
insert into musmgr.composer (full_name, biography, created_at, updated_at)
values ($1, $2, now(), now());

-- name: CreatePiece :exec
insert into musmgr.pieces (id, composed_at, description, instrumentation, title, created_at, updated_at)
values ($1, $2, $3, $4, $5, now(), now());

-- name: CreateEvent :exec
insert into musmgr.events (id, happened_at, description, event_type, create_at, updated_at)
values ($1, $2, $3, $4, now(), now());

-- name: CreatePieceEvent :exec
insert into musmgr.pieces_events (piece_id, event_id)
values ($1, $2);

-- name: CreateFile :exec
insert into musmgr.files (id, content_type, classification, file_type, name, origin, parent_id, piece_id, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6, $7, $8, now(), now());

-- name: DeletePieceEvent :exec
delete from musmgr.pieces_events
where piece_id = $1 and event_id = $2
;

-- name: DeletePiece :exec
delete from musmgr.pieces
where id = $1
;

-- name: DeleteEvent :exec
delete from musmgr.events
where id = $1
;

-- name: DeleteFile :exec
delete from musmgr.files
where id = $1
;

-- name: DeleteComposerPicture :one
update musmgr.composer set picture = null,
                           picture_content_type = null,
                           updated_at = now()
where id = true
returning picture;

-- name: UpdateComposer :exec
update musmgr.composer
set full_name = coalesce(sqlc.narg(full_name), full_name),
    biography = coalesce(sqlc.narg(biography), biography),
    updated_at = now
where id = true;

-- name: UpdateComposerPicture :exec
update musmgr.composer set picture = $1,
                           picture_content_type = $2,
                           updated_at = now()
where id = true
returning picture;

-- name: UpdatePiece :exec
update musmgr.pieces
set composed_at = coalesce(sqlc.narg(composed_at), composed_at),
    description = coalesce(sqlc.narg(description), description),
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

-- name: UpdateFileMetadata :exec
update musmgr.files
set description = coalesce(sqlc.narg(description), description),
    name = coalesce(sqlc.narg(name), name),
    updated_at = now()
where id = $1;

-- name: GetComposer :one
select *
from musmgr.composer
limit 1
;

-- name: GetEvents :many
select *
from musmgr.events
;

-- name: GetEvent :one
select *
from musmgr.events
where id = $1
;

-- name: GetPieces :many
select *
from musmgr.pieces
;

-- name: GetPiece :one
select *
from musmgr.pieces
where id = $1
;

-- name: GetPieceFiles :many
select *
from musmgr.files
where piece_id = $1 and classification <= $2
;

-- name: GetFile :one
select *
from musmgr.files
where id = $1 and classification <= $2
;

-- name: GetEventPieces :many
select id, composed_at, instrumentation, title
from musmgr.pieces
inner join musmgr.pieces_events on musmgr.pieces.id = musmgr.pieces_events.piece_id
where musmgr.pieces.id = $1
;

-- name: GetPieceEvents :many
select id, happened_at, description, event_type
from musmgr.events
inner join musmgr.pieces_events on musmgr.events.id = musmgr.pieces_events.event_id
where musmgr.pieces_events.piece_id = $1
;

-- name: GetInstrumentationNames :many
select
    unnest(enum_range(null::musmgr.instrumentation_name))::musmgr.instrumentation_name
    as instrumentation_name
;

-- name: GetEventTypes :many
select unnest(enum_range(null::musmgr.event_type))::musmgr.event_type as event_type
;

-- name: GetFileTypes :many
select unnest(enum_range(null::musmgr.file_type))::musmgr.file_type as file_type
;

