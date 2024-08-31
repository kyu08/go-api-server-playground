create table user (
  id varchar(36) not null primary key,
  screen_name varchar(20) not null unique, -- @twitter の@以降の部分
  user_name varchar(20) not null, -- `name`にしたかったがSQLのキーワードなので仕方なく`user_name`にしている
  bio varchar(160) not null,
  is_private boolean not null,
  type ENUM('public', 'private', 'deleted') not null,
  created_at timestamp not null
);

create index user_screen_name on user (screen_name);
