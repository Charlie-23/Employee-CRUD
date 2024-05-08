-- name: CreateEmployee :one
INSERT INTO employees (
  name,
  position,
  salary
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetEmployeeByID :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1;

-- name: GetEmployeeForUpdate :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateEmployee :one
UPDATE employees
SET name = $2, position = $3, salary = $4
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1;
