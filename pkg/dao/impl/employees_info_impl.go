package impl

import (
	"context"
	"demoCRUD/pkg/config"
	"demoCRUD/pkg/entity"
	"gorm.io/gorm"
)

type EmployeesInfoImpl struct {
	db *gorm.DB
}

func NewEmployeesInfoImpl() *EmployeesInfoImpl {
	return &EmployeesInfoImpl{db: config.DB}
}

func (impl EmployeesInfoImpl) AddEmployee(c context.Context, info entity.EmployeesInfo) bool {
	impl.db.Save(&info)
	return true
}

func (impl EmployeesInfoImpl) FindAllEmployee(c context.Context) []entity.EmployeesInfo {
	var infos []entity.EmployeesInfo
	impl.db.Find(&infos)
	return infos
}

func (impl EmployeesInfoImpl) UpdateEmployeeById(c context.Context, info entity.EmployeesInfo) bool {
	impl.db.Model(&entity.EmployeesInfo{}).Where("id", info.Id).Updates(entity.EmployeesInfo{Name: info.Name, Age: info.Age, Number: info.Number})
	return true
}

func (impl EmployeesInfoImpl) DeleteEmployeeById(c context.Context, id string) bool {
	impl.db.Delete(&entity.EmployeesInfo{}, id)
	return true
}
