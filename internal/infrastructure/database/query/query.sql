-- name: CreateUser :execresult
insert into user (
  id, screen_name, user_name, bio, is_private, created_at
) values (
  ?, ?, ?, ?, ?, ?
);

-- name: FindUserByScreenName :one
select * from user use index (user_screen_name)
where screen_name = ? limit 1;
