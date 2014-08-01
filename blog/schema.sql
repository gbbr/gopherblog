CREATE DATABASE IF NOT EXISTS blog;

CREATE TABLE IF NOT EXISTS posts (
	idPost int(11) NOT NULL auto_increment,
	title varchar(255),
	body blob NOT NULL,
	idUser int(11) NOT NULL,
	date timestamp,
	PRIMARY KEY (idPost)
);

CREATE TABLE IF NOT EXISTS users (
	idUser int(11) NOT NULL auto_increment,
	name varchar(255) NOT NULL,
	email varchar(100) NOT NULL,
	password char(32) NOT NULL,
	isAuthor tinyint(1),
	PRIMARY KEY (idUser)
);

CREATE TABLE IF NOT EXISTS post_tags (
	idPost int(11) NOT NULL,
	idTag int(11) NOT NULL,
	PRIMARY KEY (idPost, idTag)
);

CREATE TABLE IF NOT EXISTS tags (
	idTag int(11) NOT NULL auto_increment,
	name varchar(100) NOT NULL,
	PRIMARY KEY (idTag)
);