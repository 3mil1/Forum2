package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Storage struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("sqlite3", storage.config.DatabaseURI)

	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db

	return nil
}


func (storage *Storage) Close() {
	storage.db.Close()
}

func (storage *Storage) AddTables() {
	createSQL := `
create table  if not exists users
(
	id char(36) not null primary key ,
	email varchar(50) not null,
	login varchar(50) not null,
	password varchar(50) not null
);
create unique index  if not exists "users_email_uindex"
	on users (email);
create unique index if not exists users_login_uindex
	on users (login);
create table if not exists messages
(
	id integer not null
		constraint messages_pk
			primary key autoincrement,
	user_id char(36) not null
		references users,
	content text not null,
	created_at timestamp default CURRENT_TIMESTAMP not null,
	subject text not null default '',
	parent_id integer
		constraint messages_messages_id_fk
			references messages
);
create index if not exists messages_parent_id_index
	on messages (parent_id);
create table if not exists categories
(
	id integer not null
		constraint categories_pk
			primary key autoincrement,
	name varchar(50) not null
);
create unique index if not exists categories_name_uindex
	on categories (name);
create table if not exists "messages_categories"
(
	message_id integer not null
		constraint "messages_categories_messages_id_fk"
			references messages,
	category_id integer not null
		constraint "messages_categories_categories_id_fk"
			references categories,
	constraint "messages_categories_pk"
		primary key (message_id, category_id)
);
create table if not exists likes_dislikes
(
	message_id integer not null
		constraint likes_dislikes_messages_id_fk
			references messages,
	user_id char(36) not null
		constraint likes_dislikes_users_id_fk
			references users,
	mark boolean,
	constraint likes_dislikes_pk
		primary key (message_id, user_id)
);`
	_, err := storage.db.Exec(createSQL)
	if err != nil {
		log.Fatal(err)
	}
}
