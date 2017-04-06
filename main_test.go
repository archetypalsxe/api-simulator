package main

import (
    "io/ioutil"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHitOnRootDirectory(t *testing.T) {

    req := httptest.NewRequest("POST", "http://localhost:6000/", nil)
	w := httptest.NewRecorder()
	Index(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

    if(!strings.Contains(string(body), "Root directory of the API simulator")) {
        t.Error("Failed")
    }
}
