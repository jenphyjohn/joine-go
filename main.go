package main

import (
	"github.com/jenphyjohn/joine-go/config"
	"github.com/jenphyjohn/joine-go/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

/**
 * 初始化 iris app
 * @method NewApp
 * @return  {[type]}  api      *iris.Application  [iris app]
 */
func newApp() (app *iris.Application) {
	app = iris.New()
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(controller.ResponseResource(10404, "404 Not Found", nil))
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(controller.ResponseResource(99999, "System Error", nil))
	})

	iris.RegisterOnInterrupt(func() {

	})

	return
}

func main() {
	app := newApp()

	addr := config.GetValue("server.port")
	app.Run(iris.Addr(addr))

	//value := config.GetValue("server.contextPath")
	//fmt.Printf(value)
	//responseResult := controller.ResponseResource(10404, "404 Not Found", nil)
	//bytes, _ := json.Marshal(responseResult)
	//fmt.Print(string(bytes))
}
