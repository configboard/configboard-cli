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

func List(origin, prefix string) ([]string, error) {
	resp, err := http.Get(origin + "/api/v1/list/" + prefix)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	return strings.Split(strings.TrimSuffix(bodyStr, "\n"), "\n"), nil
}

func ListOrDie(origin, prefix string) []string {
	value, err := List(origin, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func Delete(origin, key string) error {
	resp, err := http.Get(origin + "/api/v1/delete/" + key)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	return nil
}

func DeleteOrDie(origin, key string) {
	err := Delete(origin, key)
	if err != nil {
		log.Fatal(err)
	}
}
