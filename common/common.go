package common

import (
	"Go-Blog/config"
	"Go-Blog/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

// 加载模板
func LoadTemplate(){
	// sync.WaitGroup用于同步多个goroutine
	w := sync.WaitGroup{}
	w.Add(1)
	go func(){
		// 耗时操作，需要同步
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil{
			panic(err)
		}
		w.Done()
	}()
	// 阻塞以等待go程结束
	w.Wait()
}

// 解析json数据
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Error(w http.ResponseWriter,err error){
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson,_ := json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter,data interface{}){
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson,_ := json.Marshal(result)
	w.Header().Set("Content-Type","application/json")
	_,err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}