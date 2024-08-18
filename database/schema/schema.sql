create table user (
  id varchar(36) not null primary key,
  screen_name varchar(20) not null,
  name varchar(20) not null,
  bio varchar(160) not null,
  is_private boolean not null,
  created_at timestamp not null
);
