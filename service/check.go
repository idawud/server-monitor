package service

import "net/http"

func CheckEndpintAvailabity(url string) bool {
	res, err := http.Head(url)
	if err != nil {
		return false
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true
	}
	return false
}
