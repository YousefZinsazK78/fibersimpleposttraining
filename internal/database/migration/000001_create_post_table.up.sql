CREATE TABLE IF NOT EXISTS post_tbl(
     post_id serial primary key,
     Title varchar(250) not null,
     Content varchar(300) unique not null,
     CreatedAt timestamp with time zone default CURRENT_TIMESTAMP,
     UpdatedAt timestamp with time zone default CURRENT_TIMESTAMP
);