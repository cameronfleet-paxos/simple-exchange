package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAll[T interface{}](serviceUrl string, endpoint string) ([]T, error) {
	resp, err := http.Get(serviceUrl + endpoint)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var items []T
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, err
	}
	return items, nil
}
