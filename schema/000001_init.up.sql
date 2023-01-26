CREATE TABLE usr_roles
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users
(
    id            serial primary key,
    name          varchar(255)        not null,
    username      varchar(255) unique not null,
    password_hash varchar(255)        not null,
    role_id       int references usr_roles (id) on delete cascade default 0
);

CREATE TABLE books
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255),
    rating      float        not null default 0
);

CREATE TABLE authors
(
    id         serial primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    birth      date         not null default CURRENT_DATE
);

CREATE TABLE book_author
(
    id        serial primary key,
    book_id   int references books (id) on delete cascade   not null,
    author_id int references authors (id) on delete cascade not null
);

CREATE TABLE user_book
(
    id          serial primary key,
    user_id     int references users (id) on delete cascade not null,
    book_id     int references books (id) on delete cascade not null,
    user_rating int,
    is_favorite bool
);

CREATE TABLE user_author
(
    id        serial primary key,
    user_id   int references users (id) on delete cascade   not null,
    author_id int references authors (id) on delete cascade not null
);

INSERT INTO usr_roles
VALUES
    (0, 'default_user', 'Обычный пользователь сервиса'),
    (1, 'moderator', 'Модератор сервиса'),
    (2, 'admin', 'Администратор сервиса')
