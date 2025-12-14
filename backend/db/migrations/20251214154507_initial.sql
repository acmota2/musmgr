-- +goose Up
CREATE SCHEMA IF NOT EXISTS musmgr;

CREATE type musmgr.instrumentation_name AS ENUM ('choir', 'solo', 'chamber', 'orchestra', 'opera', 'musical', 'acusmatic');

CREATE TABLE IF NOT EXISTS musmgr.works (
  id                varchar(255)         PRIMARY KEY,
  composed_at       date                 NOT NULL,
  instrumentation   musmgr.instrumentation_name NOT NULL,
  title             varchar(255)         NOT NULL
);

CREATE type musmgr.event_type AS ENUM ('concert', 'exibition', 'competition', 'festival', 'other');

CREATE TABLE IF NOT EXISTS musmgr.events (
  id varchar(255) PRIMARY KEY,
  date date NOT NULL,
  description text,
  event_type musmgr.event_type NOT NULL
);

CREATE TABLE IF NOT EXISTS musmgr.works_events (
  work_id   varchar(255) REFERENCES musmgr.works(id) NOT NULL,
  event_id  varchar(255) REFERENCES musmgr.events(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS musmgr.files (
  id varchar(255) PRIMARY KEY,
  name text NOT NULL,
  file_path text NOT NULL,
  work_id varchar(255) REFERENCES musmgr.works(id) NOT NULL
);
