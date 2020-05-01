package service

import (
	"encoding/json"
	"github.com/idawud/server-monitor/data"
	"net/http"
	"time"
)

func CheckEndpointAvailable(url string) bool {
	res, err := http.Head(url)
	if err != nil {
		return false
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true
	}
	return false
}

func GetAllAvailability() ([]byte, error) {
	var result =  make(map[string]interface{})
	for _, ep := range data.ENDPOINTS{
		result[ep] = CheckEndpointAvailable(ep)
	}
	result["timestamp"] = time.Now().String()
	return json.Marshal(result)
}