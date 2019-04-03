package initial

import (
	_ "beego_grpc/models/official_content_models"
	_ "beego_grpc/tools/upload"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/esap/wechat"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

//var Redis cache.Cache
var Bm cache.Cache

func init() {

	//初始化内存缓存
	Bm, _ = cache.NewCache("memory", `{"interval":60}`)

	//初始化日志
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	//不关闭terminal打印
	//beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)

	//初始化wechatSDK
	appID := beego.AppConfig.String("AppID")
	token := beego.AppConfig.String("Token")
	appSecret := beego.AppConfig.String("AppSecret")
	access_token := beego.AppConfig.String("Access_token")

	wd := beego.AppConfig.String("wechat_debug")

	switch {
	case wd == "true":
		wechat.Debug = true
	default:
		wechat.Debug = false
	}

	if "" == access_token {
		wechat.Set(token, appID, appSecret)
	} else {
		wechat.Set(token, appID, appSecret, access_token)
	}

	//初始化mysql数据库
	aliasName := beego.AppConfig.String("aliasName")
	drivers := beego.AppConfig.String("drivers")
	dbUserName := beego.AppConfig.String("dbUserName")
	dbPassword := beego.AppConfig.String("dbPassword")
	address := beego.AppConfig.String("address")
	dbName := beego.AppConfig.String("dbName")
	extra := beego.AppConfig.String("extra")
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	maxConn, _ := beego.AppConfig.Int("maxConn")

	err := orm.RegisterDataBase(aliasName, drivers, dbUserName+":"+dbPassword+"@tcp("+address+")/"+dbName+"?"+extra, maxIdle, maxConn)

	if nil != err {
		log.Fatalf("数据库连接失败，%v", err)

	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	//初始化redis
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("initial redis error caught: %v\n", r)
	//		Redis = nil
	//	}
	//}()
	//
	//Redis, err = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redis_host")+`"}`)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	////初始化grabc
	//var c []beego.ControllerInterface
	//c = append(c, &controllers.SiteController{}, &controllers.UserController{})
	//beego.Router("/", &controllers.SiteController{})
	//for _, v := range c {
	//	//将路由注册到beego
	//	beego.AutoRouter(v)
	//	//将路由注册到grabc
	//	grabc.RegisterController(v)
	//}
	////注册用户系统模型到grabc
	//grabc.RegisterUserModel(&models.User{})
	////增加忽律权限检查的页面
	//grabc.AppendIgnoreRoute("site", "login")
	//
	////设置grabc页面视图路径
	////如果使用默认的，不要设置或置空
	////如果需要对grabc插件进行二次开发，则需要设置这个目录，否则不需要管
	////注意：设置grabc的模板必须在beego.Run()之前设置，如果视图目录在当前项目中，可以使用相对目录，否则需要绝对路径
	//// grabc.SetViewPath("views")
	////设置grabc的layout
	//grabc.SetLayout("layout/main.html", "views")
	//
	////注册获取当前登录用户ID的函数
	//grabc.RegisterUserIdFunc(func(c *beego.Controller) int {
	//	sessionUId := c.GetSession("login_user_id")
	//
	//	if sessionUId != nil {
	//		user := models.User{}
	//		user.FindById(sessionUId.(int))
	//		return user.Id
	//	}
	//
	//	return 0
	//})

	//orm.RunCommand()
}
