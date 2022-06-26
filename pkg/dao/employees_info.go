package dao

import (
	"context"
	"demoCRUD/pkg/entity"
)

type EmployeesInfo interface {
	// 添加员工
	AddEmployee(c context.Context, info entity.EmployeesInfo) bool

	// 查找所有员工
	FindAllEmployee(c context.Context) []entity.EmployeesInfo

	// 根据id更新员工信息
	UpdateEmployeeById(c context.Context, info entity.EmployeesInfo) bool

	// 删除员工
	DeleteEmployeeById(c context.Context, id int64) bool
}
