package controller

import (
	"demoCRUD/pkg/dao/impl"
	"demoCRUD/pkg/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type EmployeesControllerImpl struct {
	dao *impl.EmployeesInfoImpl
}

type EmployeesController interface {
	AddEmployee(c *gin.Context)
	FindALlEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}

func NewEmployeesControllerImpl() *EmployeesControllerImpl {
	return &EmployeesControllerImpl{dao: impl.NewEmployeesInfoImpl()}
}

func (impl EmployeesControllerImpl) AddEmployee(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := entity.EmployeesInfo{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.AddEmployee(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "data": isOk})
}

func (impl EmployeesControllerImpl) FindAllEmployee(c *gin.Context) {
	employeesInfo := impl.dao.FindAllEmployee(c)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "data": employeesInfo})
}

func (impl EmployeesControllerImpl) UpdateEmployee(c *gin.Context) {
	//body := c.Request.Body
	//jsonBytes, err := ioutil.ReadAll(body)
	//d := json.NewDecoder(bytes.NewBuffer(jsonBytes))
	//d.UseNumber()
	//m := make(map[string]string)
	//if err != nil {
	//	panic(err)
	//}
	//info := entity.EmployeesInfo{
	//	Id:     cast.ToInt64(m["id"]),
	//	Name:   m["name"],
	//	Age:    cast.ToInt64(m["age"]),
	//	Number: m["number"],
	//}
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := entity.EmployeesInfo{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.UpdateEmployeeById(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "data": isOk})
}
func (impl EmployeesControllerImpl) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	isOk := impl.dao.DeleteEmployeeById(c, id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "data": isOk})
}
