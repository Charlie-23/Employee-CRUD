package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pzanwar/employee/util"
	"github.com/stretchr/testify/require"
)

func createRandomEmployee(t *testing.T) Employee {
	arg := CreateEmployeeParams{
		Name:     util.RandomName(),
		Position: util.RandomPosition(),
		Salary:   util.RandomSalary(),
	}

	employee, err := testStore.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	require.Equal(t, arg.Name, employee.Name)
	require.Equal(t, arg.Position, employee.Position)
	require.Equal(t, arg.Salary, employee.Salary)

	require.NotZero(t, employee.ID)
	require.NotZero(t, employee.CreatedAt)

	return employee
}

func TestCreateEmployee(t *testing.T) {
	createRandomEmployee(t)
}

func TestGetEmployee(t *testing.T) {
	employee1 := createRandomEmployee(t)
	employee2, err := testStore.GetEmployeeByID(context.Background(), employee1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, employee2)

	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, employee1.Name, employee2.Name)
	require.Equal(t, employee1.Salary, employee2.Salary)
	require.Equal(t, employee1.Position, employee2.Position)
	require.WithinDuration(t, employee1.CreatedAt.Time, employee2.CreatedAt.Time, time.Second)
}

func TestUpdateEmployee(t *testing.T) {
	employee1 := createRandomEmployee(t)

	arg := UpdateEmployeeParams{
		ID:       employee1.ID,
		Salary:   util.RandomSalary(),
		Name:     util.RandomName(),
		Position: util.RandomPosition(),
	}

	employee2, err := testStore.UpdateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee2)

	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, arg.Name, employee2.Name)
	require.Equal(t, arg.Salary, employee2.Salary)
	require.Equal(t, arg.Position, employee2.Position)
	require.WithinDuration(t, employee1.CreatedAt.Time, employee2.CreatedAt.Time, time.Second)
}

func TestDeleteEmployee(t *testing.T) {
	employee1 := createRandomEmployee(t)
	err := testStore.DeleteEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)

	employee2, err := testStore.GetEmployeeByID(context.Background(), employee1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, employee2)
}
