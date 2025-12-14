CREATE type instrumentation_name AS ENUM ('choir', 'solo', 'chamber', 'orchestra', 'opera', 'musical', 'acusmatic');

CREATE TABLE IF NOT EXISTS works (
  id                varchar(255)         PRIMARY KEY,
  composed_at       date                 NOT NULL,
  instrumentation   instrumentation_name NOT NULL,
  title             varchar(255)         NOT NULL
);

CREATE type event_type_name AS ENUM ('concert', 'exibition', 'competition', 'festival', 'other');

CREATE TABLE IF NOT EXISTS events (
  id varchar(255) PRIMARY KEY,
  date date NOT NULL,
  description text,
  event_type event_type NOT NULL
);

CREATE TABLE IF NOT EXISTS works_events (
  work_id   varchar(255) REFERENCES works(id) NOT NULL,
  event_id  varchar(255) REFERENCES events(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS files (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  file_path text NOT NULL,
  work_id varchar(255) REFERENCES works(id) NOT NULL
);
