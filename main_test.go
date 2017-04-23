package main

import (
    "io/ioutil"
    "log"
    "net/http"
    //"net/http/httptest"
    "net/url"
    "strings"
    "testing"
)

func TestThisTest(test *testing.T) {
    response, error := http.PostForm("http://localhost:6060/",
        url.Values{"key": {"Value"}, "id": {"123"}})
    if error != nil {
        log.Fatal(error)
    }

    defer response.Body.Close()
    body, error := ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(error)
    }

    if(!strings.Contains(string(body),
        "Root directory of the API simulator")) {
        test.Error("Failed")
    }
}

/*
func TestHitOnRootDirectory(test *testing.T) {
    req := httptest.NewRequest("POST", "http://localhost:6060/", nil)
    w := httptest.NewRecorder()
    Index(w, req)

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)

    if(!strings.Contains(string(body), "Root directory of the API simulator")) {
        test.Error("Failed")
    }
}
*/
