schema "public" {}

table "song" {
    schema = schema.public
    column "id" {
        type = bigserial
    }
    column "name" {
        type = varchar(255)
    }
    column "tonality_root" {
        type = smallint
    }
    column "tonality_details" {
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

table "category" {
    schema = schema.public
    column "name" {
        type = varchar(127)
    }
    column "description" {
        type = varchar(511)
    }
    primary_key {
        columns = [
            column.name
        ]
    }
    index "idx_category" {
        columns = [
            column.name
        ]
        unique = true
    }
}

table "subcategory" {
    schema = schema.public
    column "id" {
        type = bigserial
    }
    column "name" {
        type = varchar(127)
    }
    column "category_name" {
        type = varchar(127)
    }
    foreign_key "category_name" {
        columns     = [column.category_name]
        ref_columns = [table.category.column.name]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
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
    column "name" {
        type = varchar(255)
    }
    column "open" {
        type = boolean
    }
    column "type" {
        type = smallint
    }
    column "song_id" {
        type = bigint
    }
    primary_key {
        columns = [
            column.path
        ]
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
        type = bigserial
    }
    column "description" {
        type = varchar(511)
    }
    column "date" {
        type = date
    }
    column "event_type_name" {
        type = varchar(255)
    }
    foreign_key "event_type_name" {
        columns     = [column.event_type_name]
        ref_columns = [table.event_type.column.name]
        on_update   = NO_ACTION
        on_delete   = NO_ACTION
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

table "event_type" {
    schema = schema.public
    column "name" {
        type = varchar(127)
    }
    primary_key {
        columns = [
            column.name
        ]
    }
    index "idx_event_type" {
        columns = [
            column.name
        ]
        unique = true
    }
}

table "song_subcategory" {
    schema = schema.public
    column "song_id" {
        type = bigint
    }
    column "subcategory_id" {
        type = bigint
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
        type = bigint
    }
    column "song_id" {
        type = bigint
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
