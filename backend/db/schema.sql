--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.5

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
-- Name: event_type; Type: TYPE; Schema: musmgr; Owner: -
--

CREATE TYPE musmgr.event_type AS ENUM (
    'concert',
    'exibition',
    'competition',
    'festival',
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
    'acusmatic'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: events; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.events (
    id character varying(255) NOT NULL,
    date date NOT NULL,
    description text,
    event_type musmgr.event_type NOT NULL
);


--
-- Name: files; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.files (
    id character varying(255) NOT NULL,
    name text NOT NULL,
    work_id character varying(255) NOT NULL
);


--
-- Name: works; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.works (
    id character varying(255) NOT NULL,
    composed_at date NOT NULL,
    instrumentation musmgr.instrumentation_name NOT NULL,
    title character varying(255) NOT NULL
);


--
-- Name: works_events; Type: TABLE; Schema: musmgr; Owner: -
--

CREATE TABLE musmgr.works_events (
    work_id character varying(255) NOT NULL,
    event_id character varying(255) NOT NULL
);


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
-- Name: works works_pkey; Type: CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.works
    ADD CONSTRAINT works_pkey PRIMARY KEY (id);


--
-- Name: files files_work_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.files
    ADD CONSTRAINT files_work_id_fkey FOREIGN KEY (work_id) REFERENCES musmgr.works(id);


--
-- Name: works_events works_events_event_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.works_events
    ADD CONSTRAINT works_events_event_id_fkey FOREIGN KEY (event_id) REFERENCES musmgr.events(id);


--
-- Name: works_events works_events_work_id_fkey; Type: FK CONSTRAINT; Schema: musmgr; Owner: -
--

ALTER TABLE ONLY musmgr.works_events
    ADD CONSTRAINT works_events_work_id_fkey FOREIGN KEY (work_id) REFERENCES musmgr.works(id);


--
-- PostgreSQL database dump complete
--

