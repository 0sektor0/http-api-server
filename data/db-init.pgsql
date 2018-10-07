DROP  TABLE IF EXISTS vote;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS thread;
DROP TABLE IF EXISTS forum;
DROP TABLE IF EXISTS fuser;

CREATE TABLE fuser (
	id serial PRIMARY KEY,
	nickname text UNIQUE NOT NULL,
    ci_nickname text UNIQUE NOT NULL,
	email text UNIQUE NOT NULL,
    ci_email text UNIQUE NOT NULL,
	fullname text,
	about text,
	is_delited bool DEFAULT false
);

CREATE TABLE forum (
	id serial PRIMARY KEY,
	admin_id integer,
	title text,
	slug text,
	is_delited bool DEFAULT false,
	FOREIGN KEY (admin_id) REFERENCES fuser(id) ON DELETE SET NULL
);

CREATE TABLE thread (
	id serial PRIMARY KEY,
	author_id integer,
	forum_id integer,
	title text,
	message text,
	slug text,
	created timestamp,
	is_delited bool DEFAULT false,
	FOREIGN KEY (author_id) REFERENCES fuser(id) ON DELETE SET NULL,
	FOREIGN KEY (forum_id) REFERENCES forum(id) ON DELETE CASCADE
);

CREATE TABLE vote (
	id serial PRIMARY KEY,
	user_id integer NOT NULL,
	thread_id integer NOT NULL,
	voice integer default 1,
	is_delited bool DEFAULT false,
	FOREIGN KEY (user_id) REFERENCES fuser(id) ON DELETE cascade,
	FOREIGN KEY (thread_id) REFERENCES thread(id) ON DELETE CASCADE
);

CREATE TABLE post (
	id serial PRIMARY KEY,
	user_id integer,
	parent_id integer,
	thread_id integer NOT NULL,
	message text,
	edited bool DEFAULT false,
	is_delited bool DEFAULT false,
	FOREIGN KEY (user_id) REFERENCES fuser(id) ON DELETE SET NULL,
	FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE SET NULL,
	FOREIGN KEY (thread_id) REFERENCES thread(id) ON DELETE CASCADE
);