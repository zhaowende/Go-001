package impl

import (
	"Go-001/Week02/model"
	"Go-001/Week02/repository"

	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) repository.Repository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAllEmployee() ([]model.Employee, error) {
	var employeeList []model.Employee
	if err := r.db.Find(&employeeList).Error; err != nil {
		return nil, err
	}
	return employeeList, nil
}

func (r *employeeRepository) GetEmployeeById(id int) (*model.Employee, error) {
	var employee model.Employee
	if err := r.db.Where("id=?", id).First(&employee).Error; err != nil {
		// 此处对应题目内容，其他的函数的error处理尚未修改
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) SaveEmployee(employee model.Employee) error {
	if err := r.db.Save(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (r *employeeRepository) UpdateEmployee(employee model.Employee) error {
	if err := r.db.Save(&employee).Error; err != nil {
		return err
	}

	return nil
}

func (r *employeeRepository) DeleteEmployeeById(id int) error {
	var employee model.Employee
	if err := r.db.Where("id = ?", id).Delete(&employee).Error; err != nil {
		return err
	}

	return nil
}
