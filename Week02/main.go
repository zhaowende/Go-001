package main

import (
	_employeeAPI "Go-001/Week02/http"
	"Go-001/Week02/model"
	_employeeRepository "Go-001/Week02/repository/impl"
	_employeeService "Go-001/Week02/service/impl"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var httpPort string

var db *gorm.DB

func init() {
	//open a DB connection
	var err error
	db, err = gorm.Open(mysql.Open("root:root_1234@tcp(127.0.0.1:3306)/company"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database, ", err)
	}

	log.Println("db connected succeed.")
}

func main() {
	// 若数据表不存在，自动创建表
	migrate(db)

	r := gin.Default()

	employeesRouter := r.Group("/employees")
	employeeRepo := _employeeRepository.NewEmployeeRepository(db)
	employeeService := _employeeService.NewEmployeeService(employeeRepo)
	_employeeAPI.EmployeeRegister(employeesRouter, employeeService)

	_ = r.Run(":80")
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Employee{})
}
