// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: employees.sql

package db

import (
	"context"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employees (
  name,
  position,
  salary
) VALUES (
  $1, $2, $3
) RETURNING id, name, position, salary, created_at
`

type CreateEmployeeParams struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRow(ctx, createEmployee, arg.Name, arg.Position, arg.Salary)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Position,
		&i.Salary,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteEmployee, id)
	return err
}

const getEmployeeByID = `-- name: GetEmployeeByID :one
SELECT id, name, position, salary, created_at FROM employees
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEmployeeByID(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRow(ctx, getEmployeeByID, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Position,
		&i.Salary,
		&i.CreatedAt,
	)
	return i, err
}

const getEmployeeForUpdate = `-- name: GetEmployeeForUpdate :one
SELECT id, name, position, salary, created_at FROM employees
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetEmployeeForUpdate(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRow(ctx, getEmployeeForUpdate, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Position,
		&i.Salary,
		&i.CreatedAt,
	)
	return i, err
}

const updateEmployee = `-- name: UpdateEmployee :one
UPDATE employees
SET name = $2, position = $3, salary = $4
WHERE id = $1
RETURNING id, name, position, salary, created_at
`

type UpdateEmployeeParams struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employee, error) {
	row := q.db.QueryRow(ctx, updateEmployee,
		arg.ID,
		arg.Name,
		arg.Position,
		arg.Salary,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Position,
		&i.Salary,
		&i.CreatedAt,
	)
	return i, err
}