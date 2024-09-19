package functions

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseData, target)
	if err != nil {
		return err
	}
	return nil
}