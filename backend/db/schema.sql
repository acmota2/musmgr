--
-- PostgreSQL database dump
--


-- Dumped from database version 18.1 (Debian 18.1-1.pgdg13+2)
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: musmgr; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA musmgr;


--
-- Name: date_precision; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.date_precision AS ENUM (
    'day',
    'month',
    'year'
);


--
-- Name: event_type; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.event_type AS ENUM (
    'concert',
    'exhibition',
    'competition',
    'festival',
    'other'
);


--
-- Name: file_origin; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.file_origin AS ENUM (
    'user',
    'system'
);


--
-- Name: file_type; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.file_type AS ENUM (
    'score_full',
    'score_part',
    'score_preview',
    'audio_recording',
    'audio_preview',
    'picture',
    'video',
    'other'
);


--
-- Name: instrumentation_name; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.instrumentation_name AS ENUM (
    'choir',
    'solo',
    'chamber',
    'orchestra',
    'opera',
    'musical',
    'acousmatic'
);


--
-- Name: check_precision(date, musmgr.date_precision); Type: FUNCTION; Schema: musmgr; Owner: -
--

CREATE FUNCTION musmgr.check_precision(d date, p musmgr.date_precision) RETURNS boolean
    LANGUAGE sql IMMUTABLE
    AS $$
  select
    (p = 'year' and d = date_trunc('year', d)::date) or
    (p = 'month' and d = date_trunc('month', d)::date) or
    (p = 'day')
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: composer; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.composer (
    id boolean DEFAULT true NOT NULL,
    biography text NOT NULL,
    full_name text NOT NULL,
    picture uuid,
    picture_content_type text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT unique_id CHECK (id)
);


--
-- Name: events; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.events (
    id uuid NOT NULL,
    happened_at date NOT NULL,
    happened_at_precision musmgr.date_precision DEFAULT 'day'::musmgr.date_precision NOT NULL,
    description text,
    event_type musmgr.event_type NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT events_check CHECK (musmgr.check_precision(happened_at, happened_at_precision))
);


--
-- Name: files; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.files (
    id uuid NOT NULL,
    classification smallint NOT NULL,
    content_type text NOT NULL,
    description text,
    name text NOT NULL,
    origin musmgr.file_origin NOT NULL,
    file_type musmgr.file_type NOT NULL,
    parent_id uuid,
    piece_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: pieces; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.pieces (
    id uuid NOT NULL,
    composed_at date NOT NULL,
    composed_at_precision musmgr.date_precision DEFAULT 'month'::musmgr.date_precision NOT NULL,
    description text NOT NULL,
    instrumentation musmgr.instrumentation_name NOT NULL,
    title character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT pieces_check CHECK (musmgr.check_precision(composed_at, composed_at_precision))
);


--
-- Name: pieces_events; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.pieces_events (
    piece_id uuid NOT NULL,
    event_id uuid NOT NULL
);


--
-- Name: composer composer_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.composer
    ADD CONSTRAINT composer_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: pieces_events pieces_events_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.pieces_events
    ADD CONSTRAINT pieces_events_pkey PRIMARY KEY (piece_id, event_id);


--
-- Name: pieces pieces_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.pieces
    ADD CONSTRAINT pieces_pkey PRIMARY KEY (id);


--
-- Name: piece_file_access; Type: INDEX; Schema: musmgr; Owner: -
--

CREATE INDEX piece_file_access ON musmgr.files USING btree (piece_id, classification);


--
-- Name: files files_parent_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.files
    ADD CONSTRAINT files_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES musmgr.files(id) ON DELETE CASCADE;


--
-- Name: files files_piece_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.files
    ADD CONSTRAINT files_piece_id_fkey FOREIGN KEY (piece_id) REFERENCES musmgr.pieces(id) ON DELETE CASCADE;


--
-- Name: pieces_events pieces_events_event_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.pieces_events
    ADD CONSTRAINT pieces_events_event_id_fkey FOREIGN KEY (event_id) REFERENCES musmgr.events(id) ON DELETE CASCADE;


--
-- Name: pieces_events pieces_events_piece_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.pieces_events
    ADD CONSTRAINT pieces_events_piece_id_fkey FOREIGN KEY (piece_id) REFERENCES musmgr.pieces(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


