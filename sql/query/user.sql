-- name: GetUser :one
select * from user where id = ?;

-- name: CreateUser :execresult
insert into user (name) values (?);
