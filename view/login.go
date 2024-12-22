package view

import (
	"Go-Blog/common"
	"Go-Blog/config"
	"net/http"
)

func (*HtmlHandler) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}