package view

import (
	"Go-Blog/common"
	"Go-Blog/service"
	"net/http"
)

func (*HtmlHandler) Pigeonhole(w http.ResponseWriter,r *http.Request)  {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w,pigeonholeRes)
}
