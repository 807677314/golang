package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"golang.org/x/net/context"
	"io"
)

type ProgressRecord struct {
	Progresses []storage.BlkputRet `json:"progresses"`
}

func md5Hex(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func token() string {
	accessKey := beego.AppConfig.String("accessKey")
	secretKey := beego.AppConfig.String("secretKey")
	bucket := beego.AppConfig.String("bucket")

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken

}

func (this *UploadControllers) Uploadtoqiniu(io io.Reader, filename string, fpath *string, dataSize int64) error {

	upToken := token()

	//localFile := *fpath

	key := filename

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := formUploader.Put(context.Background(), &ret, upToken, key, io, dataSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//断点续传方法，存在问题未处理

	// 构建表单上传的对象
	//fileInfo, statErr := os.Stat(localFile)
	//if statErr != nil {
	//	fmt.Println(statErr)
	//	return statErr
	//}
	//fileSize := fileInfo.Size()
	//fileLmd := fileInfo.ModTime().UnixNano()
	//recordKey := md5Hex(fmt.Sprintf("%s:%s:%s:%s", bucket, key, localFile, fileLmd)) + ".progress"
	//// 指定的进度文件保存目录，实际情况下，请确保该目录存在，而且只用于记录进度文件
	//
	//recordDir := beego.AppConfig.String("recordDir")
	//mErr := os.MkdirAll(recordDir, 0755)
	//
	//if mErr != nil {
	//	fmt.Println("mkdir for record dir error,", mErr)
	//	return mErr
	//}
	//recordPath := filepath.Join(recordDir, recordKey)
	//progressRecord := ProgressRecord{}
	//// 尝试从旧的进度文件中读取进度
	//recordFp, openErr := os.Open(recordPath)
	//if openErr == nil {
	//	progressBytes, readErr := ioutil.ReadAll(recordFp)
	//	if readErr == nil {
	//		mErr := json.Unmarshal(progressBytes, &progressRecord)
	//		if mErr == nil {
	//			// 检查context 是否过期，避免701错误
	//			for _, item := range progressRecord.Progresses {
	//				if storage.IsContextExpired(item) {
	//					fmt.Println(item.ExpiredAt)
	//					progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
	//					break
	//				}
	//			}
	//		}
	//	}
	//	recordFp.Close()
	//}
	//if len(progressRecord.Progresses) == 0 {
	//	progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
	//}
	//resumeUploader := storage.NewResumeUploader(&cfg)
	//ret := storage.PutRet{}
	//progressLock := sync.RWMutex{}
	//putExtra := storage.RputExtra{
	//	Progresses: progressRecord.Progresses,
	//	Notify: func(blkIdx int, blkSize int, ret *storage.BlkputRet) {
	//		progressLock.Lock()
	//		progressLock.Unlock()
	//		//将进度序列化，然后写入文件
	//		progressRecord.Progresses[blkIdx] = *ret
	//		progressBytes, _ := json.Marshal(progressRecord)
	//		fmt.Println("write progress file", blkIdx, recordPath)
	//		wErr := ioutil.WriteFile(recordPath, progressBytes, 0644)
	//		if wErr != nil {
	//			fmt.Println("write progress file error,", wErr)
	//
	//		}
	//	},
	//}
	//err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	////上传成功之后，一定记得删除这个进度文件
	//os.Remove(recordPath)
	//fmt.Println(ret)

	return nil
}
