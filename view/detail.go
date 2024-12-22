package view

import (
	"Go-Blog/common"
	"Go-Blog/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlHandler) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取路径参数, /p/pid.html
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)

}
