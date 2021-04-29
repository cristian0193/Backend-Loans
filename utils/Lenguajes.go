package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Lenguage(lenguageCode string, messageCode string) string {
	if len(lenguageCode) <= 0 {
		lenguageCode = os.Getenv("DEFAULT_LANGUAGE")
	}
	resp, err := ioutil.ReadFile(os.Getenv(lenguageCode))
	var generic map[string]interface{}
	err = json.Unmarshal([]byte(resp), &generic)
	if err != nil {
		byteValue, _ := ioutil.ReadFile(os.Getenv("DEFAULT_LANGUAGE"))
		var generic map[string]interface{}
		json.Unmarshal([]byte(byteValue), &generic)
		return generic[messageCode].(string)
	}
	return generic[messageCode].(string)
}
