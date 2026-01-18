-- +goose Up
CREATE SCHEMA IF NOT EXISTS musmgr;

CREATE type musmgr.date_precision AS ENUM ('day', 'month', 'year');

CREATE OR REPLACE FUNCTION
  musmgr.check_precision(
    d date,
    p musmgr.date_precision
  )
RETURNS BOOLEAN
LANGUAGE sql
IMMUTABLE
AS $$
  SELECT
    (p = 'year' AND d = date_trunc('year', d)::date) OR
    (p = 'month' AND d = date_trunc('month', d)::date) OR
    (p = 'day');
$$;

CREATE type musmgr.instrumentation_name AS ENUM ('choir', 'solo', 'chamber', 'orchestra', 'opera', 'musical', 'acousmatic');

CREATE TABLE IF NOT EXISTS musmgr.pieces (
  id UUID PRIMARY KEY,
  composed_at date NOT NULL,
  composed_at_precision musmgr.date_precision NOT NULL DEFAULT 'month',
  instrumentation musmgr.instrumentation_name NOT NULL,
  title varchar(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
  CHECK(
    musmgr.date_precision(
      composed_at,
      composed_at_precision
    )
  )
);

CREATE type musmgr.event_type AS ENUM ('concert', 'exhibition', 'competition', 'festival', 'other');

CREATE TABLE IF NOT EXISTS musmgr.events (
  id UUID PRIMARY KEY,
  happened_at date NOT NULL,
  happened_at_precision musmgr.date_precision NOT NULL DEFAULT 'day',
  description text,
  event_type musmgr.event_type NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  uphappened_atd_at TIMESTAMPTZ NOT NULL DEFAULT now()
  CHECK(
    musmgr.happened_at_precision(
      happened_at,
      happened_at_precision
    )
  )
);

CREATE TABLE IF NOT EXISTS musmgr.pieces_events (
  piece_id UUID NOT NULL REFERENCES musmgr.pieces(id) ON DELETE CASCADE,
  event_id UUID NOT NULL REFERENCES musmgr.events(id) ON DELETE CASCADE,
  PRIMARY KEY (piece_id, event_id)
);

CREATE TABLE IF NOT EXISTS musmgr.files (
  id UUID PRIMARY KEY,
  name text NOT NULL,
  piece_id UUID NOT NULL REFERENCES musmgr.pieces(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
