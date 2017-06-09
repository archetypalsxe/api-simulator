package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
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

    t.Execute(self.response, context)
}

func (self *settingsPage) getApiModels() []apiModel {
    var apiModels []apiModel

    database := database{}
    database.connect()
    //database.insertData()
    rows := database.getApis()


    for rows.Next() {
        model := apiModel{}
        model.loadFromRow(rows)
        model.loadMessages()
        apiModels = append(apiModels, model)
    }

    return apiModels
}
