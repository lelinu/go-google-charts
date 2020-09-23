package services

import (
	"bytes"
	"fmt"
	"github.com/lelinu/go-google-charts/src/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

var (
	renderName      = "org-chart"
	orgChartService IOrgChartService
)

func TestMain(m *testing.M){
	orgChartService = NewOrgChartService()

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	fmt.Printf("dir is %v", dir)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestRenderToWriterValidBytes(t *testing.T){

	data, err := utils.BuildJson()
	assert.Nil(t, err)
	assert.NotNil(t, data)

	var buf bytes.Buffer
	obj := map[string]string{"jsonData": data}
	err = orgChartService.RenderToWriter(obj, renderName, &buf)

	assert.Nil(t, err)
	assert.NotNil(t, buf)
	assert.NotNil(t, buf.Bytes())
	assert.Greater(t, len(buf.Bytes()), 0)
}

