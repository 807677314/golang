package routers

import (
	"beego_grpc/controllers/official_content"
	"beego_grpc/tools/grpc_server"
	"beego_grpc/tools/upload"
	"beego_grpc/tools/wechat"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/grpc_client", &grpc_server.GrpcClientControllers{}, "GET,POST:GrpcClient")
	beego.Router("/grpc_server", &grpc_server.GrpcServerControllers{}, "GET,POST:GrpcServer")

	beego.Router("/uploadTPL", &upload.UploadControllers{}, "GET:UploadTPL")

	beego.Router("/getCategoryList", &official_content.CMSContentController{}, "GET:GetCategoryList")
	beego.Router("/getContentList", &official_content.CMSContentController{}, "POST:GetContentList")
	beego.Router("/addCategory", &official_content.CMSContentController{}, "POST:AddCategory")
	beego.Router("/getContent", &official_content.CMSContentController{}, "GET:GetContent")
	beego.Router("/addContent", &official_content.CMSContentController{}, "POST:AddContent")
	beego.Router("/removeContent", &official_content.CMSContentController{}, "POST:RemoveContent")

	beego.Router("/upload", &upload.UploadControllers{}, "POST:Upload")

	beego.Router("/wxhandler", &wechat.WechatControllers{}, "Get,Post:WxHandler")
	beego.Router("/menucreate", &wechat.WechatControllers{}, "Get,Post:MenuCreate")
	beego.Router("/menuquery", &wechat.WechatControllers{}, "Get,Post:MenuQuery")
	beego.Router("/menudelete", &wechat.WechatControllers{}, "Get,Post:MenuDelete")
	beego.Router("/menuupdate", &wechat.WechatControllers{}, "Get,Post:MenuUpdate")


}
