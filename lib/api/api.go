package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Set(origin, key, value string) error {
	_, err := http.Post(
		origin+"/api/v1/set/"+key,
		"text/plain",
		bytes.NewBuffer([]byte(value)),
	)
	if err != nil {
		return err
	}
	return nil
}

func SetOrDie(origin, key, value string) {
	err := Set(origin, key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(origin, key string) (string, error) {
	resp, err := http.Get(origin + "/api/v1/get/" + key)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyStr := string(body)
	return strings.TrimSuffix(bodyStr, "\n"), nil
}

func GetOrDie(origin, key string) string {
	value, err := Get(origin, key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
