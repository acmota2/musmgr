-- +goose Up
create schema if not exists musmgr;

create table if not exists musmgr.composer (
  id boolean primary key default true,
  biography text not null,
  full_name text not null,
  picture uuid,
  picture_content_type text,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
  constraint unique_id check(id)
);

create type musmgr.instrumentation_name as enum ('choir', 'solo', 'chamber', 'orchestra', 'opera', 'musical', 'acousmatic');

create table if not exists musmgr.pieces (
  id uuid primary key,
  composed_at text not null,
  description text not null,
  instrumentation musmgr.instrumentation_name not null,
  title varchar(255) not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create type musmgr.event_type as enum ('concert', 'exhibition', 'competition', 'festival', 'other');

create table if not exists musmgr.events (
  id uuid primary key,
  happened_at text not null,
  description text,
  event_type musmgr.event_type not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create table if not exists musmgr.pieces_events (
  piece_id uuid not null references musmgr.pieces(id) on delete cascade,
  event_id uuid not null references musmgr.events(id) on delete cascade,
  primary key (piece_id, event_id)
);

create type musmgr.file_type as enum ('score_full', 'score_part', 'score_preview', 'audio_recording', 'audio_preview', 'picture', 'video', 'other');
create type musmgr.file_origin as enum ('user', 'system');

create table if not exists musmgr.files (
  id uuid primary key,
  classification smallint not null,
  content_type text not null,
  description text,
  name text not null,
  origin musmgr.file_origin not null,
  file_type musmgr.file_type not null,
  parent_id uuid references musmgr.files(id) on delete cascade,
  piece_id uuid not null references musmgr.pieces(id) on delete cascade,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index if not exists piece_file_access
on musmgr.files (piece_id, classification);

