package upload

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/mingzhehao/goutils/filetool"
	"io"
	"strings"
	"time"
)

type UploadControllers struct {
	beego.Controller
}

func (this *UploadControllers) UploadTPL() {
	this.TplName = "upload/upload.html"

}

func (this *UploadControllers) Upload() {

	method := "news"

	if len(method) <= 0 || !uploadBizType.MatchString(method) {
		fmt.Println("upload file must input correct method. ")

		this.StopRun()
	}

	f, h, err := this.GetFile("files[]")
	if f != nil {
		f.Close()
	}

	if err != nil {
		fmt.Println("getfile err ", err)

		this.StopRun()
	} else {
		ext := filetool.Ext(h.Filename)
		fi := &FileInfo{
			Name: h.Filename,
			Type: ext,
		}

		if !fi.ValidateType() {

			this.StopRun()
		}

		var fileSize int64
		if sizeInterface, ok := f.(Sizer); ok {
			fileSize = sizeInterface.Size()
			fmt.Println(fileSize)
		}

		fileExt := strings.TrimLeft(ext, ".")
		saveName := fmt.Sprintf("%s/%s_%d%s", method, fileExt, time.Now().Unix(), ext)

		var Url string
		dataSize := h.Size
		err = this.upload(f, saveName, &Url, dataSize)

		if err != nil {
			beego.Error("upload file catch err: ", err.Error())

			this.StopRun()
		}

		GetFullResourceUrl(&Url)

		//log.Println(Url)
		this.Data["url"] = &Url
		this.StopRun()
	}
}

func (this *UploadControllers) upload(io io.Reader, saveName string, dst *string, dataSize int64) error {

	//log.Println(IsCloudStore())
	//log.Println(saveName)
	//先上传到本地，在上传到云存储服务器

	switch CloudStoreMethod() {
	//TODO 提交文件上传到 Tencent Cloud 或者提交到 AWS, 七牛云等
	case "cloud":
		return this.uploadWithCos(io, saveName, dst)
	case "qiniu":
		return this.Uploadtoqiniu(io, saveName, dst, dataSize)
	default:
		return this.uploadWithLocal(io, saveName, dst)
	}

}
