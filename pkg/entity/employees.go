package entity

type EmployeesInfo struct {
	Id     int64 `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"Age"`
	Number string `json:"Number"`
}

// 表名
func (stu EmployeesInfo) TableName() string {
	return "employees_info"
}
