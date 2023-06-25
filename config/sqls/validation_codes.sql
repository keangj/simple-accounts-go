-- name: CreateValidationCode :one
INSERT INTO validation_codes (
  email, code
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CountValidationCode :one
SELECT count(*) FROM validation_codes WHERE email = $1;
