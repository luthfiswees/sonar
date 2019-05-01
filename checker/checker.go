package checker

import (
		"encoding/json"
		"net/http"
		"io/ioutil"
		"os"
		"strings"
)

var (
	configPath = os.Getenv("GOPATH") + "/src/github.com/luthfiswees/sonar/config/config.json"
)

type Service struct {
	Name    string `json:"name"`
	BaseUrl string `json:"base-url"`
	Path    string `json:"path"`
	Port    string `json:"port"`
}

func Check() string {
	services := GetConfig()
	var serviceDetails []string

	for _, service := range services {
		_, err := http.Get(service.BaseUrl + ":" + service.Port + service.Path)
		if err != nil {
			serviceDetails = append(serviceDetails, service.Name + " is Off\n")
		} else {
			serviceDetails = append(serviceDetails, service.Name + " is On\n")
		}
	}

	return strings.Join(serviceDetails, "")
}

func GetConfig() []Service {
	jsonBytes, _ := ioutil.ReadFile(configPath)
	jsonData := string(jsonBytes)

	services := []Service{}
	json.Unmarshal([]byte(jsonData), &services)

	return services
}
