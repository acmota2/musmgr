-- name: GetSongs :many
select * from songs;

-- name: CreateSong :exec
insert into songs (id, name, tonality_root, tonality_details, subcategory_id)
values ($1, $2, $3, $4, $5);

-- name: GetEventTypes :many
select * from event_types;

-- name: CreateEvent :exec
insert into events (id, event_date, description, event_type_id)
values ($1, $2, $3, $4);

-- name: GetCategories :many
select * from categories;

-- name: GetCategorySubcategories :many
select * from subcategories
where category_id = $1;

-- name: CreateSubcategory :exec
insert into subcategories (id, name, category_id)
values ($1, $2, $3);

-- name: GetSubcategorySongs :many
select * from songs
where subcategory_id = $1;

-- name: GetSongSubcategories :many
select subcategories.id, subcategories.name, subcategories.category_id from subcategories
inner join songs
on subcategories.id = songs.subcategory_id
where songs.id = $1;

-- name: GetEventTypeEvents :many
select * from events
where event_type_id = $1;

-- name: GetSongFiles :many
select * from files
where song_id = $1;

-- name: GetEventSongs :many
select id, name, tonality_root, tonality_details, subcategory_id from songs
inner join songs_events
on songs.id = songs_events.song_id
where songs.id = $1;

-- name: CreateSongEvent :exec
insert into songs_events (song_id, event_id)
values ($1, $2);

-- name: GetTextFile :one
select * from files
where file_type = 'text' and song_id = $1;

-- name: CreateFile :exec
insert into files (id, name, file_content, file_type, song_id)
values ($1, $2, $3, $4, $5);
