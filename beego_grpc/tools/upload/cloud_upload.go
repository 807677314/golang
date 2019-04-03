package upload

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/mingzhehao/goutils/filetool"
	"github.com/nelsonken/cos-go-sdk-v5/cos"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	MIN_FILE_SIZE     = 1       // bytes
	MAX_FILE_SIZE     = 5000000 // bytes
	IMAGE_TYPES       = "(jpg|gif|p?jpeg|(x-)?png|txt)"
	ACCEPT_FILE_TYPES = IMAGE_TYPES
	UPLOAD_METHOD     = "(news|website|admin|ops|user)"
)

var (
	uploadBizType   = regexp.MustCompile(UPLOAD_METHOD)
	imageTypes      = regexp.MustCompile(IMAGE_TYPES)
	acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)
)

type FileInfo struct {
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"deleteUrl,omitempty"`
	DeleteType   string `json:"deleteType,omitempty"`
}

type Sizer interface {
	Size() int64
}

func (this *UploadControllers) getBucket() string {
	return GetConfig().Bucket
}

func (this *UploadControllers) uploadWithLocal(io io.Reader, saveName string, dstOut *string) error {

	LOCAL_FILE_DIR := beego.AppConfig.String("local_dir")

	imgPath := fmt.Sprintf("%s/%s", LOCAL_FILE_DIR, saveName)

	tmpLocalDir := LOCAL_FILE_DIR
	pos := strings.LastIndex(saveName, "/")
	if pos > 0 {
		tmpLocalDir = tmpLocalDir + "/" + saveName[0:pos]
	}

	filetool.InsureDir(tmpLocalDir)

	err := this.SaveToFile("files[]", imgPath) // 保存位置在 static/upload,没有文件夹要先创建
	if err == nil {
		*dstOut = "/" + imgPath
	}

	return err
}

func (this *UploadControllers) uploadWithCos(io io.Reader, saveName string, dstOut *string) error {

	cosClient := CosClient()

	bu := this.getBucket()
	ctx := cos.GetTimeoutCtx(time.Second * 30)

	err := cosClient.Bucket(bu).UploadObject(ctx, saveName, io, &cos.AccessControl{})

	log.Println(err)

	log.Println(err)
	if err == nil {
		*dstOut = saveName
	}

	return err
}

func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}

	fi.Error = "Filetype not allowed"
	return false
}

func (fi *FileInfo) ValidateSize() (valid bool) {
	if fi.Size < MIN_FILE_SIZE {
		fi.Error = "File is too small"
	} else if fi.Size > MAX_FILE_SIZE {
		fi.Error = "File is too big"
	} else {
		return true
	}
	return false
}
