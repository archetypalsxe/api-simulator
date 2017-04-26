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

func TestSettingsPage(test *testing.T) {
    sendRequestForm(test,
        "http://localhost:6060/settings",
        url.Values{},
        "This is the settings page")
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

func TestWorldspanPricingResponse(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<BPC9>"),
        "BPW9")
}

func TestWorldspanCardAuthorization(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<HOS_CMD>CK/"),
        "OK - APVL CODE IS")
}

func TestWorldspanUpdatePnr(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<UPC7>"),
        "UPW7")
}

func TestWorldspanDisplayPnrNative(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<HOS_CMD>*"),
        "<HOS_RSP_SCR>M</HOS_RSP_SCR>")
}

func TestWorldspanTicketing(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<HOS_CMD>EZEI#$*"),
        "TKT NBR")
}

func TestWorldspanCloseSession(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<HOS_RSP_SCR>F</HOS_RSP_SCR>"),
        "<RSP_COU>0000</RSP_COU>")
}

func TestWorldspanDisplayPnr(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("<DPC8>"),
        "DPW8")
}

func TestWorldspanInvalidRequest(test *testing.T) {
    sendRequestPost(test,
        "http://localhost:6060/worldspan",
        strings.NewReader("Invalid Request"),
        "Type of request not found")
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
