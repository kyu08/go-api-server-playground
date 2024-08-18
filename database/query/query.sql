-- -- name: GetAuthor :one
-- select * from authors
-- where id = ? limit 1;
--
-- -- name: ListAuthors :many
-- select * from authors
-- order by name2;
--
-- -- name: CreateAuthor :execresult
-- insert into authors (
--   name2, bio
-- ) values (
--   ?, ?
-- );
--
-- -- name: DeleteAuthor :exec
-- delete from authors
-- where id = ?;

-- name: CreateUser :execresult
insert into user (
  id, screen_name, user_name,
  bio, is_private, created_at
) values (
  ?, ?, ?, ?, ?, ?
);

-- name: GetUserByScreenName :one
select * from user
where screen_name = ? limit 1;
