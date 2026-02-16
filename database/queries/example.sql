-- name: GetExampleByID :one
SELECT * FROM examples WHERE id = ? LIMIT 1;