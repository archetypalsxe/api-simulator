package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"
    "strconv"
);

type settingsPage struct {
    response http.ResponseWriter
}

func (self *settingsPage) respond() {
    dataPath := os.Getenv("GOPATH") + "/src/api-simulator/htmlTemplates/"
    t, error := template.ParseFiles(dataPath + "settings.html")
    if error != nil {
        log.Fatal(error)
    }
    t.Execute(self.response, map[string] string {"Title": "Whoa"})

    self.getData()
}

func (self *settingsPage) getData() {
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
        fmt.Fprintln(self.response,
            strconv.Itoa(id) + " " + name + " " + beginningEscape +
            " " + endingEscape)
    }

    var template string
    rows = database.getResponses()
    for rows.Next() {
        rows.Scan(&id, &template)
        fmt.Fprintln(self.response,
            strconv.Itoa(id) + " " + template)
    }


    var apiId int
    var identifier string
    var responseId int
    rows = database.getMessages()
    for rows.Next() {
        rows.Scan(&id, &apiId, &identifier, &responseId)
        fmt.Fprintln(self.response, strconv.Itoa(id) +
            " " + strconv.Itoa(apiId) + " " + identifier + " " +
            strconv.Itoa(responseId))
    }
}
