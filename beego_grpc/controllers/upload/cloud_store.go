package upload

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nelsonken/cos-go-sdk-v5/cos"
)

type CloudStoreConfig struct {
	AppID     string
	SecretID  string
	SecretKey string
	Region    string
	Domain    string
	Bucket    string
}

func (s *CloudStoreConfig) ConvertDown(op *cos.Option) {
	s.AppID = op.AppID
	s.SecretID = op.SecretID
	s.SecretKey = op.SecretKey
	s.Region = op.Region
	s.Domain = op.Domain
	s.Bucket = op.Bucket
}

func (s *CloudStoreConfig) ConvertUp(op *cos.Option) {
	op.AppID = s.AppID
	op.SecretID = s.SecretID
	op.SecretKey = s.SecretKey
	op.Region = s.Region
	op.Domain = s.Domain
	op.Bucket = s.Bucket
}

var (
	storeConfig *CloudStoreConfig
	cosClient   *cos.Client
	cloudStore  string
)

func init() {
	var err error
	cloudStore = beego.AppConfig.String("cloud_store")

	if err != nil {
		beego.Error("can't find cloud configuration!")
		return
	}

	appId := beego.AppConfig.String("cloud_app_id")
	secretId := beego.AppConfig.String("cloud_secret_id")
	secretKey := beego.AppConfig.String("cloud_secret_key")
	region := beego.AppConfig.String("cloud_region")
	domain := beego.AppConfig.String("cloud_domain")
	bucket := beego.AppConfig.String("cloud_bucket")

	storeConfig = &CloudStoreConfig{
		AppID:     appId,
		SecretID:  secretId,
		SecretKey: secretKey,
		Region:    region,
		Domain:    domain,
		Bucket:    bucket,
	}

	if "local" == cloudStore {
		beego.Info("disable cloud store!")
		return
	} else {
		op := cos.Option{}
		storeConfig.ConvertUp(&op)
		cosClient = cos.New(&op)
	}
}

func GetConfig() *CloudStoreConfig {
	if storeConfig == nil {
		storeConfig = &CloudStoreConfig{}
	}

	return storeConfig
}

// 腾讯云存储Client
func CosClient() *cos.Client {
	return cosClient
}

func cloudServerHost() string {
	protocol := "https"
	return fmt.Sprintf("%s://%s-%s.cos.%s.%s/", protocol, storeConfig.Bucket, storeConfig.AppID, storeConfig.Region, storeConfig.Domain)
}

func GetFullResourceUrl(uri *string) string {
	if "cloud" == cloudStore {
		*uri = cloudServerHost() + *uri
	} else {

	}

	return *uri
}

func CloudStoreMethod() string {
	return cloudStore
}
