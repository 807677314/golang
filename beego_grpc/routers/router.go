package routers

import (
	"beego_grpc/controllers/official_content"
	"beego_grpc/controllers/upload"
	"beego_grpc/tools/grpc_server"
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

	beego.Router("/upload",&upload.UploadControllers{},"POST:Upload")

}
