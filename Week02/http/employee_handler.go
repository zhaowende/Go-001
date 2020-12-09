package http

import (
	"Go-001/Week02/model"
	"Go-001/Week02/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service service.Service
}

func EmployeeRegister(router *gin.RouterGroup, service service.Service) {
	handler := EmployeeHandler{
		service: service,
	}

	router.GET("/", handler.getAllEmployee)
	router.GET("/:id", handler.getEmployee)
	router.POST("/", handler.addEmployee)
	router.PUT("/:id", handler.updateEmployee)
	router.DELETE("/:id", handler.deleteEmployee)
}

func (h *EmployeeHandler) getAllEmployee(c *gin.Context) {
	employees, err := h.service.GetAllEmployee()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) getEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	employee, err := h.service.GetEmployeeById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) addEmployee(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.SaveEmployee(employee)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *EmployeeHandler) updateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateEmployee(employee)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *EmployeeHandler) deleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteEmployeeById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nil)
}
