-- name: FindUserByID :one
select * from user
where id = ? limit 1;
