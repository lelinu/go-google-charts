package services

import (
	"bytes"
	"github.com/lelinu/go-google-charts/src/templates"
	"html/template"
	"io"
)

type OrgChartService struct {
}

type IOrgChartService interface {
	RenderToWriter(data interface{}, renderName string, w ...io.Writer) error
}

func NewOrgChartService() IOrgChartService {
	return &OrgChartService{}
}

//renderChart will load the related golang template
func renderChart(data interface{}, w io.Writer, name string) error {
	contents := []string{templates.OrgChartTpl}

	tpl := template.Must(template.New("").Parse(contents[0]))
	return tpl.ExecuteTemplate(w, name, data)
}

//RenderToWriter will translate the template outcome to bytes
func (s *OrgChartService) RenderToWriter(data interface{}, renderName string, w ...io.Writer) error {
	var b bytes.Buffer
	if err := renderChart(data, &b, renderName); err != nil {
		return err
	}

	for i := 0; i < len(w); i++ {
		_, err := w[i].Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}