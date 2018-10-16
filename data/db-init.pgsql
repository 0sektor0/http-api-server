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
	FOREIGN KEY (author_id) REFERENCES fuser(id) ON DELETE CASCADE,
	FOREIGN KEY (forum_id) REFERENCES forum(id) ON DELETE CASCADE
);

CREATE TABLE vote (
	id 			serial PRIMARY KEY,
	user_id 	integer NOT NULL,
	thread_id 	integer NOT NULL,
	voice 		integer default 1,
	is_delited 	bool DEFAULT false,
	UNIQUE (user_id, thread_id),
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
	path     	integer [],
	FOREIGN KEY (user_id) REFERENCES fuser(id) ON DELETE SET NULL,
	FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE SET NULL,
	FOREIGN KEY (thread_id) REFERENCES thread(id) ON DELETE CASCADE
);


DROP INDEX IF EXISTS index_on_forum_slug;
CREATE UNIQUE INDEX  index_on_forum_slug
ON forum (ci_slug);


DROP INDEX IF EXISTS index_on_user_email;
CREATE UNIQUE INDEX index_on_user_email
ON fuser (ci_email);


DROP INDEX IF EXISTS index_on_user_nickname;
CREATE UNIQUE INDEX index_on_user_nickname
ON fuser (ci_nickname);


DROP INDEX IF EXISTS index_on_user_nickname_and_email;
CREATE UNIQUE INDEX index_on_user_nickname_and_email
ON fuser (ci_nickname, ci_email);


DROP INDEX IF EXISTS index_on_post_thread;
CREATE INDEX index_on_post_thread
ON post (thread_id);


DROP INDEX IF EXISTS index_on_post_parent;
CREATE INDEX index_on_post_parent
ON post (parent_id);


DROP INDEX IF EXISTS index_on_post_path;
CREATE INDEX index_on_post_path
ON post USING GIN (path);


DROP INDEX IF EXISTS index_on_post_path_and_thread;
CREATE INDEX index_on_post_path_and_thread
ON post (thread_id, parent_id);


DROP INDEX IF EXISTS index_on_thread_slug;
CREATE UNIQUE INDEX index_on_thread_slug 
ON thread (ci_slug) 
WHERE ci_slug != '';


DROP INDEX IF EXISTS index_on_thread_forum;
CREATE INDEX index_on_thread_forum 
ON thread (forum_id);


DROP INDEX IF EXISTS index_on_thread_slug_and_created;
CREATE INDEX index_on_thread_slug_and_created 
ON thread (ci_slug, created);


--CREATE TABLE IF NOT EXISTS "members" (
--  forum  CITEXT,
--  author CITEXT
--);


--DROP INDEX IF EXISTS index_on_vote_user_and_thread;
--CREATE UNIQUE INDEX index_on_vote_user_and_thread 
--ON vote (thread, "user");


--DROP INDEX IF EXISTS index_on_member_forum_and_author;
--CREATE UNIQUE INDEX index_on_member_forum_and_author 
--ON members (forum_id, author);


--DROP INDEX IF EXISTS index_on_members_forum;
--CREATE INDEX index_on_members_forum ON members (forum);
