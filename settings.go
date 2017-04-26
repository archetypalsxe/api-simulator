package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
);

type settingsPage struct {
    request *http.Request;
    response http.ResponseWriter;
}

func (self *settingsPage) respond() {
    dataPath := os.Getenv("GOPATH") + "/htmlTemplates/";
    t, error := template.ParseFiles(dataPath + "settings.html")
    if error != nil {
        log.Fatal(error)
    }
    t.Execute(self.response, map[string] string {"Title": "Whoa"})
}
