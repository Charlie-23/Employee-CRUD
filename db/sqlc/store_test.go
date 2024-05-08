package db

import (
	"context"
	"testing"

	"github.com/pzanwar/employee/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEmployeeTx(t *testing.T) {
	arg := CreateEmployeeParams{
		Name:     util.RandomName(),
		Position: util.RandomPosition(),
		Salary:   util.RandomSalary(),
	}
	request := CreateEmployeeTxParams{
		CreateEmployeeParams: arg,
	}


	result, err := testStore.CreateEmployeeTx(context.Background(), request)
	result.Employee.ID = 1
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.Name, result.Employee.Name)
	require.Equal(t, arg.Position, result.Employee.Position)
	require.Equal(t, arg.Salary, result.Employee.Salary)

	require.NotZero(t, result.Employee.ID)
	require.NotZero(t, result.Employee.CreatedAt)

}
