package view

import (
	"Go-Blog/common"
	"Go-Blog/service"
	"net/http"
)

func (*HtmlHandler) Writing(w http.ResponseWriter, r *http.Request)  {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w,wr)
}
