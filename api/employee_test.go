package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/pzanwar/employee/db/mock"
	db "github.com/pzanwar/employee/db/sqlc"
	"github.com/pzanwar/employee/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEmployeeAPI(t *testing.T) {
	employee := randomEmployee()

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":     employee.Name,
				"position": employee.Position,
				"salary":   employee.Salary,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateEmployeeParams{
					Name:     employee.Name,
					Position: employee.Position,
					Salary:   employee.Salary,
				}

				store.EXPECT().
					CreateEmployee(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, employee)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":     employee.Name,
				"position": employee.Position,
				"salary":   employee.Salary,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateEmployee(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Employee{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/employees"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdateEmployeeAPI(t *testing.T) {
	employee := randomEmployee()

	updateName := util.RandomName()

	updatedEmployee := db.Employee{
		ID:       employee.ID,
		Name:     updateName,
		Salary:   employee.Salary,
		Position: employee.Position,
	}

	testCases := []struct {
		name          string
		body          gin.H
		empID         int64
		setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK Update Name",
			body: gin.H{
				"name": updateName,
			},
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateEmployeeParams{
					ID:       employee.ID,
					Name:     updateName,
					Position: employee.Position,
					Salary:   employee.Salary,
				}

				gomock.InOrder(
					store.EXPECT().
						GetEmployeeByID(gomock.Any(), gomock.Eq(employee.ID)).
						Times(1).
						Return(employee, nil),

					store.EXPECT().
						UpdateEmployee(gomock.Any(), gomock.Eq(arg)).
						Times(1).
						Return(updatedEmployee, nil),
				)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, updatedEmployee)
			},
		},
		{
			name: "NotFound",
			body: gin.H{
				"name": employee.Name,
			},
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetEmployeeByID(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(db.Employee{}, db.ErrRecordNotFound)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name": employee.Name,
			},
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetEmployeeByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Employee{}, sql.ErrConnDone)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/update/%d", tc.empID)
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetEmployeeAPI(t *testing.T) {
	employee := randomEmployee()

	testCases := []struct {
		name          string
		empID         int64
		setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByID(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, employee)
			},
		},
		{
			name:  "NotFound",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Employee{}, db.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployeeByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Employee{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/employees/%d", tc.empID)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeleteEmployeeAPI(t *testing.T) {
	employee := randomEmployee()

	testCases := []struct {
		name          string
		empID         int64
		setupAuth     func(t *testing.T, request *http.Request)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEmployee(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "NotFound",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEmployee(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			empID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteEmployee(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/delete/%d", tc.empID)

			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomEmployee() db.Employee {
	return db.Employee{
		ID:       util.RandomInt(10, 1000),
		Name:     util.RandomName(),
		Salary:   util.RandomSalary(),
		Position: util.RandomPosition(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, employee db.Employee) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotEmployee db.Employee
	err = json.Unmarshal(data, &gotEmployee)
	require.NoError(t, err)
	require.Equal(t, employee, gotEmployee)
}
