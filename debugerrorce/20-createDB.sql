--
-- PostgreSQL database dump
--

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


CREATE DATABASE roman WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.UTF-8';


ALTER DATABASE roman OWNER TO spostgres;

\connect roman

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
SET default_table_access_method = heap;


CREATE TABLE public.users (
    id serial primary key,
    name text not null,
    comment text
);

ALTER TABLE public.users OWNER TO spostgres;


CREATE TABLE public.countries (
    id serial primary key,
    name text NOT NULL,
    comment text
);

ALTER TABLE public.countries OWNER TO spostgres;


CREATE TABLE public.tribunes (
   id serial primary key,
   country_id integer,
   user_id integer,
   foreign key(country_id) references public.countries(id),
   foreign key(user_id) references public.users(id)
);

ALTER TABLE public.tribunes OWNER TO spostgres;

-- =================================================================================================


COPY public.users (id, name, comment) FROM stdin WITH csv;
"1","Gaius Julius Caesar",NULL
"2","Augustus",NULL
"3","Tiberius",NULL
"4","Calligula",NULL
"5","Claudius",NULL
"6","Nero",NULL
\.


COPY public.countries (id, name, comment) FROM stdin WITH csv;
"41","Switzerland",NULL
"44","United Kingdom",NULL
"49","German",NULL
\.

COPY public.tribunes (id, country_id, user_id) FROM stdin WITH delimiter ',' NULL AS 'NULL' csv;
"1","41","1"
"2","41","6"
"3","49","4"
\.

-- =======================================================================================================================

REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT USAGE on SCHEMA public to PUBLIC;

GRANT SELECT                                                   on table public.users to psqltestro;
GRANT SELECT                                                   ON TABLE public.countries TO psqltestro;
GRANT SELECT                                                   ON TABLE public.tribunes TO psqltestro;

GRANT SELECT,INSERT,UPDATE                                     on table public.users to psqltestrw;
GRANT SELECT,INSERT,UPDATE                                     ON TABLE public.countries TO psqltestrw;
GRANT SELECT,INSERT,UPDATE                                     ON TABLE public.tribunes TO psqltestrw;

GRANT SELECT,INSERT,UPDATE,DELETE                              on table public.users to psqltestrwd;
GRANT SELECT,INSERT,UPDATE,DELETE                              ON TABLE public.countries TO psqltestrwd;
GRANT SELECT,INSERT,UPDATE,DELETE                              ON TABLE public.tribunes TO psqltestrwd;

GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE                     on table public.users to psqltestrwdt;
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE                     ON TABLE public.countries TO psqltestrwdt;
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE                     ON TABLE public.tribunes TO psqltestrwdt;

-- EOF

