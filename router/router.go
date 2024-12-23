package router

import (
	"Go-Blog/api"
	"Go-Blog/view"
	"net/http"
)

// 注册url的回调
// 共三种类型：
// 页面：view
// 接口：Api
// 静态资源：

func Router(){
	// 页面
	// index
	http.HandleFunc("/", view.Html.Index)
	// category
	http.HandleFunc("/c/", view.Html.Category)
	// login
	http.HandleFunc("/login",view.Html.Login)
	// detail
	http.HandleFunc("/p/", view.Html.Detail)
	// writting 
	http.HandleFunc("/writing", view.Html.Writing)
	// pigeonhole
	http.HandleFunc("/pigeonhole", view.Html.Pigeonhole)

	// 接口
	// 将/resource/js/index.js这样的请求路径映射到/public对应的文件夹下
	http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
	// 登录
	http.HandleFunc("/api/v1/login", api.Api.Login)
	// 发布或编辑文章
	http.HandleFunc("/api/v1/post", api.Api.SaveAndUpdatePost)
	// 请求文章
	http.HandleFunc("/api/v1/post/",api.Api.GetPost)
	// 搜索
	http.HandleFunc("/api/v1/post/search",api.Api.SearchPost)
	// 上传图片
	http.HandleFunc("/api/v1/qiniu/token", api.Api.QiniuToken)
}