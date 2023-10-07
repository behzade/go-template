-- name: GetUser :one
select * from localdb.user where id = ?;

-- name: GetUsers :many
select * from localdb.user limit ? offset ?;

-- name: CreateUser :execresult
insert into localdb.user (name) values (?);
