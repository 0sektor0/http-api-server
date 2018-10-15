DROP  TABLE IF EXISTS vote;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS thread;
DROP TABLE IF EXISTS forum;
DROP TABLE IF EXISTS fuser;

CREATE TABLE fuser (
	id 			serial PRIMARY KEY,
	nickname 	text UNIQUE NOT NULL,
    ci_nickname text UNIQUE NOT NULL,
	email 		text UNIQUE NOT NULL,
    ci_email 	text UNIQUE NOT NULL,
	fullname 	text,
	about 		text,
	is_delited 	bool DEFAULT false
);

CREATE TABLE forum (
	id 			serial PRIMARY KEY,
	admin_id 	integer NOT NULL,
	title 		text NOT NULL,
	slug 		text UNIQUE NOT NULL,
	ci_slug 	text UNIQUE NOT NULL,
	is_delited 	bool DEFAULT false,
	FOREIGN KEY (admin_id) REFERENCES fuser(id)
);

CREATE TABLE thread (
	id 			serial PRIMARY KEY,
	author_id 	integer NOT NULL,
	forum_id 	integer NOT NULL,
	title 		text,
	message 	text,
	slug 		text UNIQUE,
	ci_slug 	text UNIQUE,
	created 	timestamp with time zone,
	is_delited 	bool DEFAULT false,
	UNIQUE (author_id, forum_id, title, message, ci_slug, created),
	FOREIGN KEY (author_id) REFERENCES fuser(id) ON DELETE CASCADE,
	FOREIGN KEY (forum_id) REFERENCES forum(id) ON DELETE CASCADE
);

CREATE TABLE vote (
	id 			serial PRIMARY KEY,
	user_id 	integer NOT NULL,
	thread_id 	integer NOT NULL,
	voice 		integer default 1,
	is_delited 	bool DEFAULT false,
	FOREIGN KEY (user_id) REFERENCES fuser(id) ON DELETE CASCADE,
	FOREIGN KEY (thread_id) REFERENCES thread(id) ON DELETE CASCADE
);

CREATE TABLE post (
	id 			serial PRIMARY KEY,
	user_id 	integer NOT NULL,
	thread_id 	integer NOT NULL,
	parent_id 	integer,
	message 	text,
	edited 		bool DEFAULT false,
	is_delited 	bool DEFAULT false,
	created 	timestamp with time zone,
	FOREIGN KEY (user_id) REFERENCES fuser(id) ON DELETE SET NULL,
	FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE SET NULL,
	FOREIGN KEY (thread_id) REFERENCES thread(id) ON DELETE CASCADE
);