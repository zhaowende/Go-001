package impl

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllEmployee(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:root_1234@tcp(127.0.0.1:3306)/company"), &gorm.Config{})
	_ = err
	fmt.Println("db:", db)
	println("i am a test!")

	repo := NewEmployeeRepository(db)
	fmt.Printf("repo: %v", repo)

	list, err := repo.GetAllEmployee()
	if err != nil {
		fmt.Printf("error\n: %v", err)
	}

	fmt.Printf("result:%v\n", list)
}
