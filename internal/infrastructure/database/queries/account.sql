-- name: FindAccountById :one
select
  *
from
  accounts a
where
  a.id = $1 Limit 1;