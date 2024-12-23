package api

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"Go-Blog/common"
	"Go-Blog/config"
	"net/http"
)

func (*ApiHandler) QiniuToken(w http.ResponseWriter,r *http.Request)  {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "cat-go-blog-server"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.Success(w,upToken)
}

