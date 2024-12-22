package api

import (
	"Go-Blog/common"
	"Go-Blog/models"
	"Go-Blog/service"
	"Go-Blog/utils"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*ApiHandler) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	//获取用户id，判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期，请重新登录"))
		return
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	// post请求，发布文章
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid: -1,
			Title: title,
			Slug: slug,
			Content: content,
			Markdown: markdown,
			CategoryId: categoryId,
			UserId: uid,
			ViewCount: 0,
			Type: pType,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// put请求，编辑文章
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			Pid: pid,
			Title: title,
			Slug: slug,
			Content: content,
			Markdown: markdown,
			CategoryId: categoryId,
			UserId: uid,
			ViewCount: 0,
			Type: pType,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*ApiHandler) GetPost(w http.ResponseWriter,r *http.Request)  {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path,"/api/v1/post/")
	pid,err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w,errors.New("不识别此请求路径"))
		return
	}
	post,err := service.GetPostByID(pid)
	if err != nil {
		common.Error(w,err)
		return
	}
	common.Success(w,post)
}

func (*ApiHandler) SearchPost(w http.ResponseWriter,r *http.Request)  {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w,searchResp)
}