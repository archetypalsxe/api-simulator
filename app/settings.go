package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    //"strconv"
);

type settingsPage struct {
    response http.ResponseWriter
}

type SettingsContext struct {
    ApiModels []apiModel
}


func (self *settingsPage) respond() {
    dataPath := os.Getenv("GOPATH") + "/src/api-simulator/htmlTemplates/"
    t, error := template.ParseFiles(dataPath + "settings.html")
    if error != nil {
        log.Fatal(error)
    }

    apiModels := self.getApiModels()
    var context SettingsContext
    context = SettingsContext{ApiModels: apiModels}

    for i := 0; i < 1; i++ {
        fmt.Fprintln(self.response, apiModels[i].name)
    }

    t.Execute(self.response, context)
}

func (self *settingsPage) getApiModels() []apiModel {
    var apiModels []apiModel

    database := database{}
    database.connect()
    database.insertData()
    rows := database.getApis()


    var id int
    var name string
    var beginningEscape string
    var endingEscape string
    for rows.Next() {
        rows.Scan(&id, &name, &beginningEscape, &endingEscape)
        apiModels = append(apiModels, apiModel{id: id, name: name,
                beginningEscape:beginningEscape,
                endingEscape:endingEscape})
        }

    return apiModels
}
