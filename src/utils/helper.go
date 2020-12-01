package utils

import (
	"encoding/json"
	"github.com/lelinu/go-google-charts/src/models"
	"io/ioutil"
)

var (
	filePath = "json/test.json"
)

func BuildJson() (string, error) {

	// load file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// validate with model
	var rows []models.Row
	err = json.Unmarshal(data, &rows)
	if err != nil{
		return "", err
	}

	return string(data), nil
}
