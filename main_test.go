package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
    "testing"
)

func TestRootDirectory(test *testing.T) {
    sendRequest(test,
        "http://localhost:6060",
        url.Values{},
        "Root directory of the API simulator")
}

func TestExampleId(test *testing.T) {
    sendRequest(test,
        "http://localhost:6060/example/5",
        url.Values{},
        "Provided ID:  5")
}
        //url.Values{"key": {"Value"}, "id": {"123"}})

func sendRequest(test *testing.T,
    url string,
    urlValues url.Values,
    expectedText string) {
    response, error := http.PostForm(url, urlValues)
    if error != nil {
        log.Fatal(error)
    }

    defer response.Body.Close()
    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(error)
    }

    if(!strings.Contains(string(body),
        expectedText)) {
            test.Error("Did not find expected text: '" + expectedText +
            "' Received: '"+ strings.TrimSpace(string(body)) + "'")
    }
}
