package main

import (
	"github.com/lelinu/go-google-charts/src/services"
	"github.com/lelinu/go-google-charts/src/utils"
	"net/http"
)

var (
	orgChartService services.IOrgChartService
)

func init(){
	orgChartService = services.NewOrgChartService()
}

func orgChartHandler(w http.ResponseWriter, _ *http.Request) {

	data, err := utils.BuildJson()
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	obj := map[string]string{"jsonData": data}
	orgChartService.RenderToWriter(obj, "org-chart", w)
}

func gaugeChartHandler(w http.ResponseWriter, _ *http.Request){

	obj := map[string]float64{"riskAssessmentValue": 1.0}
	orgChartService.RenderToWriter(obj, "gauge-chart", w)

}

func main(){
	http.HandleFunc("/", orgChartHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
