package http

import (
	"encoding/json"
	"errors"
	"gin-web/model"
	"gin-web/module/employee/mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var engine *gin.Engine
var mockEmployeeService *mocks.Service

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mockEmployeeService = new(mocks.Service)

	gin.SetMode(gin.TestMode)
	engine = gin.New()
	v1Group := engine.Group("v1")
	employeeRouter := v1Group.Group("/employees")
	EmployeeRegister(employeeRouter, mockEmployeeService)
}

func TestEmployeeHandler_getAllEmployee(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		employeeList := make([]model.Employee, 0)
		employeeList = append(employeeList, model.Employee{
			BaseModel: model.BaseModel{
				ID: 1,
			},
			Name:      "zzz",
			Address:   "yyy",
			Telephone: "15900236644",
			Email:     "zzs@gmail.com",
		})
		mockEmployeeService.On("GetAllEmployee").Return(employeeList, nil).Once()

		request, _ := http.NewRequest(http.MethodGet, "/v1/employees/", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusOK, writerRecorder.Code)
		respBody, _ := ioutil.ReadAll(writerRecorder.Body)
		var employeeListResp []model.Employee
		json.Unmarshal(respBody, &employeeListResp)
		assert.Equal(t, employeeList, employeeListResp)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEmployeeService.On("GetAllEmployee").Return(nil, errors.New("db error")).Once()

		request, _ := http.NewRequest(http.MethodGet, "/v1/employees/", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestEmployeeHandler_getEmployee(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		em := &model.Employee{
			BaseModel: model.BaseModel{
				ID: 1,
			},
			Name:      "zzz",
			Address:   "yyy",
			Telephone: "15900236644",
			Email:     "zzs@gmail.com",
		}
		mockEmployeeService.On("GetEmployeeById", mock.Anything).Return(em, nil).Once()

		request, _ := http.NewRequest(http.MethodGet, "/v1/employees/1", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)
		assert.Equal(t, http.StatusOK, writerRecorder.Code)

		respBody, _ := ioutil.ReadAll(writerRecorder.Body)
		var employeeResp *model.Employee
		json.Unmarshal(respBody, &employeeResp)
		assert.Equal(t, em, employeeResp)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockEmployeeService.On("GetEmployeeById", mock.Anything).Return(nil, errors.New("data not found")).Once()

		request, _ := http.NewRequest(http.MethodGet, "/v1/employees/1", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusNotFound, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestEmployeeHandler_addEmployee(t *testing.T) {
	employeeJson := `
						{
							"name": "abc",
							"address": "edf",
							"telephone": "15900237777",
							"email":"abc@gmail.com"
						}`

	t.Run("success", func(t *testing.T) {
		mockEmployeeService.On("SaveEmployee", mock.Anything).Return(nil).Once()

		request, _ := http.NewRequest(http.MethodPost, "/v1/employees/", strings.NewReader(employeeJson))
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)
		assert.Equal(t, http.StatusCreated, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("client-error-failed", func(t *testing.T) {
		mockEmployeeService.On("SaveEmployee", mock.Anything).Return(errors.New("client parameter error")).Maybe()

		request, _ := http.NewRequest(http.MethodPost, "/v1/employees/", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusBadRequest, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("db-error-failed", func(t *testing.T) {
		mockEmployeeService.On("SaveEmployee", mock.Anything).Return(errors.New("db error")).Maybe()

		request, _ := http.NewRequest(http.MethodPost, "/v1/employees/", strings.NewReader(employeeJson))
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestEmployeeHandler_updateEmployee(t *testing.T) {
	employeeJson := `
						{
							"id": "1",
							"name": "abc",
							"address": "edf",
							"telephone": "15900237777",
							"email":"abc@gmail.com"
						}`

	t.Run("success", func(t *testing.T) {
		mockEmployeeService.On("UpdateEmployee", mock.Anything).Return(nil).Once()

		request, _ := http.NewRequest(http.MethodPut, "/v1/employees/1", strings.NewReader(employeeJson))
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)
		assert.Equal(t, http.StatusOK, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("client-error-failed", func(t *testing.T) {
		mockEmployeeService.On("UpdateEmployee", mock.Anything).Return(errors.New("client parameter error")).Maybe()

		request, _ := http.NewRequest(http.MethodPut, "/v1/employees/1", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusBadRequest, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("db-error-failed", func(t *testing.T) {
		mockEmployeeService.On("UpdateEmployee", mock.Anything).Return(errors.New("db error")).Maybe()

		request, _ := http.NewRequest(http.MethodPut, "/v1/employees/1", strings.NewReader(employeeJson))
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}

func TestEmployeeHandler_deleteEmployee(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockEmployeeService.On("DeleteEmployeeById", mock.Anything).Return(nil).Once()

		request, _ := http.NewRequest(http.MethodDelete, "/v1/employees/1", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)
		assert.Equal(t, http.StatusOK, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})

	t.Run("db-error-failed", func(t *testing.T) {
		mockEmployeeService.On("DeleteEmployeeById", mock.Anything).Return(errors.New("db error")).Maybe()

		request, _ := http.NewRequest(http.MethodDelete, "/v1/employees/1", nil)
		writerRecorder := httptest.NewRecorder()
		engine.ServeHTTP(writerRecorder, request)

		assert.Equal(t, http.StatusInternalServerError, writerRecorder.Code)

		mockEmployeeService.AssertExpectations(t)
	})
}
