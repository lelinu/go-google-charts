package main

import (
	"github.com/lelinu/go-google-charts/src/services"
	"github.com/lelinu/go-google-charts/src/utils"
	"net/http"
)

var (
	renderName      = "org-chart"
	orgChartService services.IOrgChartService
)

func init(){
	orgChartService = services.NewOrgChartService()
}

func handler(w http.ResponseWriter, _ *http.Request) {

	data, err := utils.BuildJson()
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	obj := map[string]string{"jsonData": data}
	orgChartService.RenderToWriter(obj, renderName, w)
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
