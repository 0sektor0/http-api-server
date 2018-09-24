CREATE TABLE Users (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	nickname varchar(20) NOT NULL,
	email varchar(50) NOT NULL,
	fullname varchar(50),
	is_delited bool DEFAULT false,
	about varchar(255)
);

CREATE TABLE Forums (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	admin_id INTEGER,
	title VARCHAR(50) DEFAULT "",
	slug VARCHAR(255) DEFAULT "",
	is_delited bool DEFAULT false,
	FOREIGN KEY (admin_id) REFERENCES Users(id) ON DELETE SET NULL
);

CREATE TABLE Threads (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	author_id INTEGER,
	forum_id INTEGER,
	title VARCHAR(100) DEFAULT "",
	message varchar(500) DEFAULT "",
	slug VARCHAR(255) default "",
	created datetime,
	is_delited bool DEFAULT false,
	FOREIGN KEY (author_id) REFERENCES Users(id) ON DELETE SET NULL,
	FOREIGN KEY (forum_id) REFERENCES Forums(id) ON DELETE CASCADE
);

CREATE TABLE Votes (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	user_id INTEGER NOT NULL,
	thread_id INTEGER NOT NULL,
	voice tinyint default 1,
	is_delited bool DEFAULT false,
	FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE cascade,
	FOREIGN KEY (thread_id) REFERENCES Threads(id) ON DELETE CASCADE
);

CREATE TABLE Posts (
	id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	user_id INTEGER,
	parent_id INTEGER,
	thread_id INTEGER NOT NULL,
	message varchar(500) default "",
	edited bool DEFAULT false,
	is_delited bool DEFAULT false,
	FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE SET NULL,
	FOREIGN KEY (parent_id) REFERENCES Posts(id) ON DELETE SET NULL,
	FOREIGN KEY (thread_id) REFERENCES Threads(id) ON DELETE CASCADE
);