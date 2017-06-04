package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func doDELETE(url string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return body
}

func doPOSTWithBytes(url string, xBytes []byte) []byte {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(xBytes))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return body
}

func doPOST(url string, jsonString string) []byte {
	jsonBytes := []byte(jsonString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func doPOSTObject(url string, x interface{}) []byte {
	xBytes, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return doPOSTWithBytes(url, xBytes)
}

func doGET(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return body
}
