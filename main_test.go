package main

import (
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
    "testing"
)

func TestRootDirectory(test *testing.T) {
    sendRequestForm(test,
        "http://localhost:6060",
        url.Values{},
        "Root directory of the API simulator")
}

func TestExampleId(test *testing.T) {
    sendRequestForm(test,
        "http://localhost:6060/example/5",
        url.Values{},
        "Provided ID:  5")
}

func TestWorldspanPowerShopper(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<PSC5>"),
        "PSW5")
}
        //url.Values{"key": {"Value"}, "id": {"123"}})

func sendRequestForm(test *testing.T,
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

    validateResponse(test, strings.TrimSpace(string(body)), expectedText)

}

func sendRequestPost(test *testing.T,
    url string,
    data io.Reader,
    expectedText string) {

    request, error := http.NewRequest("POST",
        url,
        data)
    if error != nil {
        log.Fatal(error)
    }

    httpClient := http.Client{}
    response, error := httpClient.Do(request)
    if error != nil {
        log.Fatal(error)
    }
    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(error)
    }
    validateResponse(test, strings.TrimSpace(string(body)), expectedText)
}

func validateResponse(test *testing.T, actual string, expected string) {
    if(!strings.Contains(actual,
        expected)) {
            test.Error("Did not find expected text: '" + expected +
            "' Received: '"+ actual + "'")
    }
}
