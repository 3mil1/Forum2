create table if not exists users
(
    id       INTEGER     not null PRIMARY KEY AUTOINCREMENT,
    email    varchar(50) not null,
    login    varchar(50) not null,
    password varchar(50) not null
);
create unique index if not exists "users_email_uindex"
    on users (email);
create unique index if not exists users_login_uindex
    on users (login);
create table if not exists messages
(
    id         integer  not null
        constraint messages_pk
            primary key autoincrement,
    user_id    char(36) not null
        references users,
    content    text     not null,
    created_at timestamp         default CURRENT_TIMESTAMP not null,
    subject    text     not null default '',
    parent_id  integer
        constraint messages_messages_id_fk
            references messages
);
create index if not exists messages_parent_id_index
    on messages (parent_id);
create table if not exists categories
(
    id   integer     not null
        constraint categories_pk
            primary key autoincrement,
    name varchar(50) not null
);
create unique index if not exists categories_name_uindex
    on categories (name);
create table if not exists "messages_categories"
(
    message_id  integer not null
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
    message_id integer  not null
        constraint likes_dislikes_messages_id_fk
            references messages,
    user_id    char(36) not null
        constraint likes_dislikes_users_id_fk
            references users,
    mark       boolean,
    constraint likes_dislikes_pk
        primary key (message_id, user_id)
);