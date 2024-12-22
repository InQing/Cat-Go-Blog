package view

import (
	"Go-Blog/common"
	"Go-Blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}

// 处理index页面
func (*HtmlHandler) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	if err := r.ParseForm();err != nil{
		log.Println("表单获取失败：",err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page,_ = strconv.Atoi(pageStr)
	}
	pageSize := 10

	path := r.URL.Path
	slug := strings.TrimPrefix(path,"/")

	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}

	index.WriteData(w, hr)
}
