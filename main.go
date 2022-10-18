/*
data:2022-10-10
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"os"
	"time"
	"webProject/controller"
	"webProject/datasource"
	"webProject/service"
)

func main() {
	//datasource.Engine_sql()
	app := NewApp()

	//配置文件
	Configuration(app)

	//读取服务器配置json文件
	config := InitAppConfig()

	addr := "localhost:" + config.Port
	//路由设置处理
	mvcHandle(app)
	//启动iris服务
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //序列化更快
	)
}

func NewApp() *iris.Application {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")

	//注册静态资源,教程中的staticweb方法不能使用
	app.HandleDir("/static", "./static")
	app.HandleDir("/img", "./static/img")
	app.HandleDir("/manage/static", "./static")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	//处理根目录get访问,这部分也不能使用func(cxt context.Context)使用下文
	app.Get("/", func(cxt iris.Context) {
		cxt.View("index.html")
	})

	return app
}

// 服务器配置iris
func Configuration(app *iris.Application) {
	//app.Configure(iris.WithConfiguration(iris.Configuration{
	//		Charset: "UTF-8",
	//	}))
	//配置字符
	app.Configure(iris.WithCharset("UTF_8"))

	//错误配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
		ctx.View("errors/404.html")
	})

	app.OnErrorCode(iris.StatusServiceUnavailable, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"statusCode": 500,
			"massage":    "service error",
		})
	})
}

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

func InitAppConfig() *AppConfig {
	//打开文件
	file, err := os.Open("./config.json")
	HandleError(err)
	defer file.Close()
	//读取数据
	//jsonData, err := ioutil.ReadAll()
	decoder := json.NewDecoder(file)
	con := AppConfig{}
	err = decoder.Decode(&con)
	HandleError(err)
	return &con
}

// 错误处理

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// mvc架构处理
func mvcHandle(app *iris.Application) {
	//启用Session,主要用于设置登录超时
	fmt.Println("-------------mvc handle---------")
	sessionManage := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	//获取数据库引擎
	engine := datasource.Engine_sql()
	//获取数据服务
	adminService := service.NewAdminService(engine)

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessionManage.Start,
	)
	//MVC处理各种get\post等，在controller中写
	admin.Handle(new(controller.AdminController))

	//统计功能
	statisServer := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}/"))
	statis.Register(
		statisServer,
		sessionManage.Start,
	)
	statis.Handle(new(controller.StatisController))
}
