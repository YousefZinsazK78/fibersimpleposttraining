CREATE TABLE IF NOT EXISTS user_tbl(
     user_id serial primary key,
     Username varchar(200) not null,
     Email varchar(300) unique not null,
     HashPassword text not null,
     CreatedAt timestamp with time zone default CURRENT_TIMESTAMP,
     UpdatedAt timestamp with time zone default CURRENT_TIMESTAMP
);