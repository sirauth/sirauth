--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2

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

--
-- Name: sirauth; Type: DATABASE; Schema: -; Owner: sirauth
--

CREATE DATABASE sirauth WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE sirauth OWNER TO sirauth;

\connect sirauth

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

--
-- Name: template_keys; Type: TYPE; Schema: public; Owner: sirauth
--

CREATE TYPE public.template_keys AS ENUM (
    'sign_up',
    'sign_in',
    'sign_out',
    'user_info'
);


ALTER TYPE public.template_keys OWNER TO sirauth;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: application_domains; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.application_domains (
    id integer NOT NULL,
    application_id uuid NOT NULL,
    domain_id uuid NOT NULL
);


ALTER TABLE public.application_domains OWNER TO sirauth;

--
-- Name: application_domains_id_seq; Type: SEQUENCE; Schema: public; Owner: sirauth
--

CREATE SEQUENCE public.application_domains_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.application_domains_id_seq OWNER TO sirauth;

--
-- Name: application_domains_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sirauth
--

ALTER SEQUENCE public.application_domains_id_seq OWNED BY public.application_domains.id;


--
-- Name: application_templates; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.application_templates (
    id uuid NOT NULL,
    application_id uuid NOT NULL,
    kind public.template_keys NOT NULL,
    domain_ids integer[] NOT NULL
);


ALTER TABLE public.application_templates OWNER TO sirauth;

--
-- Name: application_users; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.application_users (
    id integer NOT NULL,
    application_id uuid NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.application_users OWNER TO sirauth;

--
-- Name: application_users_id_seq; Type: SEQUENCE; Schema: public; Owner: sirauth
--

CREATE SEQUENCE public.application_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.application_users_id_seq OWNER TO sirauth;

--
-- Name: application_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sirauth
--

ALTER SEQUENCE public.application_users_id_seq OWNED BY public.application_users.id;


--
-- Name: applications; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.applications (
    id uuid NOT NULL,
    name character varying(200) NOT NULL
);


ALTER TABLE public.applications OWNER TO sirauth;

--
-- Name: domains; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.domains (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(200) NOT NULL
);


ALTER TABLE public.domains OWNER TO sirauth;

--
-- Name: users; Type: TABLE; Schema: public; Owner: sirauth
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    email character varying(240) NOT NULL,
    password character varying(150) NOT NULL
);


ALTER TABLE public.users OWNER TO sirauth;

--
-- Name: application_domains id; Type: DEFAULT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_domains ALTER COLUMN id SET DEFAULT nextval('public.application_domains_id_seq'::regclass);


--
-- Name: application_users id; Type: DEFAULT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_users ALTER COLUMN id SET DEFAULT nextval('public.application_users_id_seq'::regclass);


--
-- Data for Name: application_domains; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.application_domains (id, application_id, domain_id) FROM stdin;
\.


--
-- Data for Name: application_templates; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.application_templates (id, application_id, kind, domain_ids) FROM stdin;
\.


--
-- Data for Name: application_users; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.application_users (id, application_id, user_id) FROM stdin;
\.


--
-- Data for Name: applications; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.applications (id, name) FROM stdin;
\.


--
-- Data for Name: domains; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.domains (id, name) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: sirauth
--

COPY public.users (id, email, password) FROM stdin;
\.


--
-- Name: application_domains_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sirauth
--

SELECT pg_catalog.setval('public.application_domains_id_seq', 1, false);


--
-- Name: application_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sirauth
--

SELECT pg_catalog.setval('public.application_users_id_seq', 1, false);


--
-- Name: application_domains application_domains_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_domains
    ADD CONSTRAINT application_domains_pk PRIMARY KEY (id);


--
-- Name: application_templates application_templates_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_templates
    ADD CONSTRAINT application_templates_pk PRIMARY KEY (id);


--
-- Name: application_users application_users_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_users
    ADD CONSTRAINT application_users_pk PRIMARY KEY (id);


--
-- Name: applications applications_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.applications
    ADD CONSTRAINT applications_pk PRIMARY KEY (id);


--
-- Name: domains domains_names; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.domains
    ADD CONSTRAINT domains_names UNIQUE (name);


--
-- Name: domains domains_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.domains
    ADD CONSTRAINT domains_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: application_domains application_domains_applications_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_domains
    ADD CONSTRAINT application_domains_applications_id_fk FOREIGN KEY (application_id) REFERENCES public.applications(id);


--
-- Name: application_domains application_domains_domains_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_domains
    ADD CONSTRAINT application_domains_domains_id_fk FOREIGN KEY (domain_id) REFERENCES public.domains(id);


--
-- Name: application_users application_users_applications_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_users
    ADD CONSTRAINT application_users_applications_id_fk FOREIGN KEY (application_id) REFERENCES public.applications(id);


--
-- Name: application_users application_users_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: sirauth
--

ALTER TABLE ONLY public.application_users
    ADD CONSTRAINT application_users_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

