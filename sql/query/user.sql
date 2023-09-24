-- name: GetUser :one
select * from localdb.user where id = ?;

-- name: CreateUser :execresult
insert into localdb.user (name) values (?);
