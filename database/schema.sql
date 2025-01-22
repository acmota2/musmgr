CREATE TABLE IF NOT EXISTS categories (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  description text
);

CREATE TABLE IF NOT EXISTS subcategories (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  category_id varchar(255) REFERENCES categories(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS songs (
  id                varchar(255) PRIMARY KEY,
  name              varchar(255)    NOT NULL,
  tonality_root     int             NOT NULL,
  tonality_details  varchar(63)     NOT NULL,
  subcategory_id    varchar(255) REFERENCES subcategories(id) NOT NULL
);


CREATE TABLE IF NOT EXISTS event_types (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  description text
);

CREATE TABLE IF NOT EXISTS events (
  id varchar(255) PRIMARY KEY,
  event_date date NOT NULL,
  description text,
  event_type_id varchar(255) REFERENCES event_types(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS songs_events (
  song_id   varchar(255) REFERENCES songs(id) NOT NULL,
  event_id  varchar(255) REFERENCES events(id) NOT NULL
);

CREATE TYPE file_type AS ENUM ('text', 'score', 'song');

CREATE TABLE IF NOT EXISTS files (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  file_path text NOT NULL,
  file_type file_type NOT NULL,
  song_id varchar(255) REFERENCES songs(id) NOT NULL
);

