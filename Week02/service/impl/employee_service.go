package impl

import (
	"Go-001/Week02/model"
	"Go-001/Week02/repository"
	"Go-001/Week02/service"
)

type employeeService struct {
	repo repository.Repository
}

func NewEmployeeService(repo repository.Repository) service.Service {
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) GetAllEmployee() ([]model.Employee, error) {
	return s.repo.GetAllEmployee()
}

func (s *employeeService) GetEmployeeById(id int) (*model.Employee, error) {
	return s.repo.GetEmployeeById(id)
}

func (s *employeeService) SaveEmployee(employee model.Employee) error {
	return s.repo.SaveEmployee(employee)
}

func (s *employeeService) UpdateEmployee(employee model.Employee) error {
	return s.repo.UpdateEmployee(employee)
}

func (s *employeeService) DeleteEmployeeById(id int) error {
	return s.repo.DeleteEmployeeById(id)
}
