schema "public" {}

table "song" {
    schema = schema.public
    column "id" {
        type = int
    }
    column "name" {
        type = varchar(255)
    }
    column "tonality" {
        type = varchar(24)
    }
    primary_key {
        columns = [
            column.id
        ]
    }
    index "idx_song" {
        columns = [
            column.id
        ]
        unique = true
    }
}

table "subcategory" {
    schema = schema.public
    column "id" {
        type = int
    }
    column "name" {
        type = varchar(255)
    }
    primary_key {
        columns = [
            column.id
        ]
    }
    index "idx_subcategory" {
        columns = [
            column.id
        ]
        unique = true
    }
}

table "file" {
    schema = schema.public
    column "path" {
        type = varchar(255)
    }
    column "type" {
        type = smallint
    }
    column "name" {
        type = varchar(255)
    }
    primary_key {
        columns = [
            column.path
        ]
    }
    column "song_id" {
        type = int
    }
    foreign_key "song_id" {
        columns     = [column.song_id]
        ref_columns = [table.song.column.id]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
    }
    index "idx_file" {
        columns = [
            column.path
        ]
        unique = true
    }
}

table "event" {
    schema = schema.public
    column "id" {
        type = int
    }
    column "name" {
        type = varchar(255)
    }
    column "date" {
        type = date
    }
    primary_key {
        columns = [
            column.id
        ]
    }
    index "idx_event" {
        columns = [
            column.id
        ]
        unique = true
    }
}

table "song_subcategory" {
    schema = schema.public
    column "song_id" {
        type = int
    }
    column "subcategory_id" {
        type = int
    }
    primary_key {
        columns = [
            column.song_id,
            column.subcategory_id
        ]
    }
    foreign_key "song_id" {
        columns     = [column.song_id]
        ref_columns = [table.song.column.id]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
    }
    foreign_key "subcategory_id" {
        columns     = [column.subcategory_id]
        ref_columns = [table.subcategory.column.id]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
    }
    index "idx_song_subcategory" {
        columns = [
            column.song_id,
            column.subcategory_id
        ]
        unique = true
    }
}

table "song_event" {
    schema = schema.public
    column "event_id" {
        type = int
    }
    column "song_id" {
        type = int
    }
    foreign_key "event_id" {
        columns     = [column.event_id]
        ref_columns = [table.subcategory.column.id]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
    }
    foreign_key "song_id" {
        columns     = [column.song_id]
        ref_columns = [table.song.column.id]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
    }
    index "idx_song_event" {
        columns = [
            column.song_id,
            column.event_id
        ]
        unique = true
    }
}
