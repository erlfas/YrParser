package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func doPOSTWithBytes(url string, xBytes []byte) []byte {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(xBytes))
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return body
}

func doPOST(url string, x interface{}) []byte {
	xBytes, err := json.MarshalIndent(x, "", " ")
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return doPOSTWithBytes(url, xBytes)
}

func doGET(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err.Error())
		panic(err)
	}

	return body
}
