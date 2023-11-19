CREATE TABLE IF NOT EXISTS post_tbl(
     post_id serial primary key,
     username varchar(150) unique not null,
     password varchar(250) not null,
     email varchar(250) unique not null,
     created_at timestamp with time zone default CURRENT_TIMESTAMP,
     updated_at timestamp with time zone default CURRENT_TIMESTAMP,
);