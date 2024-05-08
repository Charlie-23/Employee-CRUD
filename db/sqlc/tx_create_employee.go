package db

import "context"

type CreateEmployeeTxParams struct {
	CreateEmployeeParams
}

type CreateEmployeeTxResult struct {
	Employee Employee
}

func (store *SQLStore) CreateEmployeeTx(ctx context.Context, arg CreateEmployeeTxParams) (CreateEmployeeTxResult, error) {
	var result CreateEmployeeTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Employee, err = q.CreateEmployee(ctx, arg.CreateEmployeeParams)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
