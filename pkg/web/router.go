package web

import (
	"demoCRUD/pkg/config"
	"demoCRUD/pkg/web/controller"
	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	// 拦截器
	r.Use(config.HttpInterceptor())
	// 跨域
	r.Use(config.CorsConfig())
	// 路由组
	routerGroup := r.Group("/")
	{
		routerGroup.POST("/add/", controller.NewEmployeesControllerImpl().AddEmployee)
		routerGroup.GET("/find/", controller.NewEmployeesControllerImpl().FindAllEmployee)
		routerGroup.POST("/update/", controller.NewEmployeesControllerImpl().UpdateEmployee)
		routerGroup.POST("/delete/:id", controller.NewEmployeesControllerImpl().DeleteEmployee)
	}

	r.Run("127.0.0.1:" + config.PORT)
}
