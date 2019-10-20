--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: events; Type: TABLE; Schema: public; Owner: calendar_user
--

CREATE TABLE public.events (
    id uuid NOT NULL,
    user_name text NOT NULL,
    title text NOT NULL,
    note text,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL
);


ALTER TABLE public.events OWNER TO calendar_user;

--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: calendar_user
--

COPY public.events (id, user_name, title, note, start_time, end_time) FROM stdin;
\.


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: calendar_user
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: owner_idx; Type: INDEX; Schema: public; Owner: calendar_user
--

CREATE INDEX owner_idx ON public.events USING btree (user_name);


--
-- PostgreSQL database dump complete
--

