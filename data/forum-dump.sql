--
-- PostgreSQL database dump
--

-- Dumped from database version 10.5 (Debian 10.5-1.pgdg90+1)
-- Dumped by pg_dump version 10.5 (Debian 10.5-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: forum; Type: TABLE; Schema: public; Owner: forum_admin
--

CREATE TABLE public.forum (
    id integer NOT NULL,
    admin_id integer NOT NULL,
    title text NOT NULL,
    slug text NOT NULL,
    ci_slug text NOT NULL,
    is_delited boolean DEFAULT false
);


ALTER TABLE public.forum OWNER TO forum_admin;

--
-- Name: forum_id_seq; Type: SEQUENCE; Schema: public; Owner: forum_admin
--

CREATE SEQUENCE public.forum_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.forum_id_seq OWNER TO forum_admin;

--
-- Name: forum_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: forum_admin
--

ALTER SEQUENCE public.forum_id_seq OWNED BY public.forum.id;


--
-- Name: fuser; Type: TABLE; Schema: public; Owner: forum_admin
--

CREATE TABLE public.fuser (
    id integer NOT NULL,
    nickname text NOT NULL,
    ci_nickname text NOT NULL,
    email text NOT NULL,
    ci_email text NOT NULL,
    fullname text,
    about text,
    is_delited boolean DEFAULT false
);


ALTER TABLE public.fuser OWNER TO forum_admin;

--
-- Name: fuser_id_seq; Type: SEQUENCE; Schema: public; Owner: forum_admin
--

CREATE SEQUENCE public.fuser_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fuser_id_seq OWNER TO forum_admin;

--
-- Name: fuser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: forum_admin
--

ALTER SEQUENCE public.fuser_id_seq OWNED BY public.fuser.id;


--
-- Name: post; Type: TABLE; Schema: public; Owner: forum_admin
--

CREATE TABLE public.post (
    id integer NOT NULL,
    user_id integer NOT NULL,
    thread_id integer NOT NULL,
    parent_id integer,
    message text,
    edited boolean DEFAULT false,
    is_delited boolean DEFAULT false,
    created timestamp with time zone,
    path integer[]
);


ALTER TABLE public.post OWNER TO forum_admin;

--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: forum_admin
--

CREATE SEQUENCE public.post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO forum_admin;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: forum_admin
--

ALTER SEQUENCE public.post_id_seq OWNED BY public.post.id;


--
-- Name: thread; Type: TABLE; Schema: public; Owner: forum_admin
--

CREATE TABLE public.thread (
    id integer NOT NULL,
    author_id integer NOT NULL,
    forum_id integer NOT NULL,
    title text,
    message text,
    slug text,
    ci_slug text,
    created timestamp with time zone,
    is_delited boolean DEFAULT false
);


ALTER TABLE public.thread OWNER TO forum_admin;

--
-- Name: thread_id_seq; Type: SEQUENCE; Schema: public; Owner: forum_admin
--

CREATE SEQUENCE public.thread_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.thread_id_seq OWNER TO forum_admin;

--
-- Name: thread_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: forum_admin
--

ALTER SEQUENCE public.thread_id_seq OWNED BY public.thread.id;


--
-- Name: vote; Type: TABLE; Schema: public; Owner: forum_admin
--

CREATE TABLE public.vote (
    id integer NOT NULL,
    user_id integer NOT NULL,
    thread_id integer NOT NULL,
    voice integer DEFAULT 1,
    is_delited boolean DEFAULT false
);


ALTER TABLE public.vote OWNER TO forum_admin;

--
-- Name: vote_id_seq; Type: SEQUENCE; Schema: public; Owner: forum_admin
--

CREATE SEQUENCE public.vote_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vote_id_seq OWNER TO forum_admin;

--
-- Name: vote_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: forum_admin
--

ALTER SEQUENCE public.vote_id_seq OWNED BY public.vote.id;


--
-- Name: forum id; Type: DEFAULT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.forum ALTER COLUMN id SET DEFAULT nextval('public.forum_id_seq'::regclass);


--
-- Name: fuser id; Type: DEFAULT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser ALTER COLUMN id SET DEFAULT nextval('public.fuser_id_seq'::regclass);


--
-- Name: post id; Type: DEFAULT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);


--
-- Name: thread id; Type: DEFAULT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread ALTER COLUMN id SET DEFAULT nextval('public.thread_id_seq'::regclass);


--
-- Name: vote id; Type: DEFAULT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.vote ALTER COLUMN id SET DEFAULT nextval('public.vote_id_seq'::regclass);


--
-- Data for Name: forum; Type: TABLE DATA; Schema: public; Owner: forum_admin
--

COPY public.forum (id, admin_id, title, slug, ci_slug, is_delited) FROM stdin;
\.


--
-- Data for Name: fuser; Type: TABLE DATA; Schema: public; Owner: forum_admin
--

COPY public.fuser (id, nickname, ci_nickname, email, ci_email, fullname, about, is_delited) FROM stdin;
\.


--
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: forum_admin
--

COPY public.post (id, user_id, thread_id, parent_id, message, edited, is_delited, created, path) FROM stdin;
\.


--
-- Data for Name: thread; Type: TABLE DATA; Schema: public; Owner: forum_admin
--

COPY public.thread (id, author_id, forum_id, title, message, slug, ci_slug, created, is_delited) FROM stdin;
\.


--
-- Data for Name: vote; Type: TABLE DATA; Schema: public; Owner: forum_admin
--

COPY public.vote (id, user_id, thread_id, voice, is_delited) FROM stdin;
\.


--
-- Name: forum_id_seq; Type: SEQUENCE SET; Schema: public; Owner: forum_admin
--

SELECT pg_catalog.setval('public.forum_id_seq', 1, false);


--
-- Name: fuser_id_seq; Type: SEQUENCE SET; Schema: public; Owner: forum_admin
--

SELECT pg_catalog.setval('public.fuser_id_seq', 1, false);


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: forum_admin
--

SELECT pg_catalog.setval('public.post_id_seq', 1, false);


--
-- Name: thread_id_seq; Type: SEQUENCE SET; Schema: public; Owner: forum_admin
--

SELECT pg_catalog.setval('public.thread_id_seq', 1, false);


--
-- Name: vote_id_seq; Type: SEQUENCE SET; Schema: public; Owner: forum_admin
--

SELECT pg_catalog.setval('public.vote_id_seq', 1, false);


--
-- Name: forum forum_ci_slug_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.forum
    ADD CONSTRAINT forum_ci_slug_key UNIQUE (ci_slug);


--
-- Name: forum forum_pkey; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.forum
    ADD CONSTRAINT forum_pkey PRIMARY KEY (id);


--
-- Name: forum forum_slug_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.forum
    ADD CONSTRAINT forum_slug_key UNIQUE (slug);


--
-- Name: fuser fuser_ci_email_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser
    ADD CONSTRAINT fuser_ci_email_key UNIQUE (ci_email);


--
-- Name: fuser fuser_ci_nickname_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser
    ADD CONSTRAINT fuser_ci_nickname_key UNIQUE (ci_nickname);


--
-- Name: fuser fuser_email_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser
    ADD CONSTRAINT fuser_email_key UNIQUE (email);


--
-- Name: fuser fuser_nickname_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser
    ADD CONSTRAINT fuser_nickname_key UNIQUE (nickname);


--
-- Name: fuser fuser_pkey; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.fuser
    ADD CONSTRAINT fuser_pkey PRIMARY KEY (id);


--
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- Name: thread thread_ci_slug_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread
    ADD CONSTRAINT thread_ci_slug_key UNIQUE (ci_slug);


--
-- Name: thread thread_pkey; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread
    ADD CONSTRAINT thread_pkey PRIMARY KEY (id);


--
-- Name: thread thread_slug_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread
    ADD CONSTRAINT thread_slug_key UNIQUE (slug);


--
-- Name: vote vote_pkey; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.vote
    ADD CONSTRAINT vote_pkey PRIMARY KEY (id);


--
-- Name: vote vote_user_id_thread_id_key; Type: CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.vote
    ADD CONSTRAINT vote_user_id_thread_id_key UNIQUE (user_id, thread_id);


--
-- Name: index_on_forum_slug; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE UNIQUE INDEX index_on_forum_slug ON public.forum USING btree (ci_slug);


--
-- Name: index_on_post_parent; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_post_parent ON public.post USING btree (parent_id);


--
-- Name: index_on_post_path; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_post_path ON public.post USING gin (path);


--
-- Name: index_on_post_path_and_thread; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_post_path_and_thread ON public.post USING btree (thread_id, parent_id);


--
-- Name: index_on_post_thread; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_post_thread ON public.post USING btree (thread_id);


--
-- Name: index_on_thread_forum; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_thread_forum ON public.thread USING btree (forum_id);


--
-- Name: index_on_thread_slug; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE UNIQUE INDEX index_on_thread_slug ON public.thread USING btree (ci_slug) WHERE (ci_slug <> ''::text);


--
-- Name: index_on_thread_slug_and_created; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE INDEX index_on_thread_slug_and_created ON public.thread USING btree (ci_slug, created);


--
-- Name: index_on_user_email; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE UNIQUE INDEX index_on_user_email ON public.fuser USING btree (ci_email);


--
-- Name: index_on_user_nickname; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE UNIQUE INDEX index_on_user_nickname ON public.fuser USING btree (ci_nickname);


--
-- Name: index_on_user_nickname_and_email; Type: INDEX; Schema: public; Owner: forum_admin
--

CREATE UNIQUE INDEX index_on_user_nickname_and_email ON public.fuser USING btree (ci_nickname, ci_email);


--
-- Name: forum forum_admin_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.forum
    ADD CONSTRAINT forum_admin_id_fkey FOREIGN KEY (admin_id) REFERENCES public.fuser(id);


--
-- Name: post post_parent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.post(id) ON DELETE SET NULL;


--
-- Name: post post_thread_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_thread_id_fkey FOREIGN KEY (thread_id) REFERENCES public.thread(id) ON DELETE CASCADE;


--
-- Name: post post_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.fuser(id) ON DELETE SET NULL;


--
-- Name: thread thread_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread
    ADD CONSTRAINT thread_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.fuser(id) ON DELETE CASCADE;


--
-- Name: thread thread_forum_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.thread
    ADD CONSTRAINT thread_forum_id_fkey FOREIGN KEY (forum_id) REFERENCES public.forum(id) ON DELETE CASCADE;


--
-- Name: vote vote_thread_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.vote
    ADD CONSTRAINT vote_thread_id_fkey FOREIGN KEY (thread_id) REFERENCES public.thread(id) ON DELETE CASCADE;


--
-- Name: vote vote_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: forum_admin
--

ALTER TABLE ONLY public.vote
    ADD CONSTRAINT vote_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.fuser(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

