package main

import (
    "fmt"
    "html"
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
    dataPath := os.Getenv("GOPATH") + "/htmlTemplates/"
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
    for rows.Next() {
        rows.Scan(&id, &name)
        fmt.Fprintln(self.response, strconv.Itoa(id) + ": " + name);
    }

    var apiId int
    var identifier string
    var responseId int
    rows = database.getApis()
log.Print(rows)
    for rows.Next() {
        rows.Scan(&id, &apiId, &identifier, &responseId)
log.Print(identifier)
        fmt.Fprintln(self.response, html.EscapeString(strconv.Itoa(id) +
            " " + strconv.Itoa(apiId) + " " + identifier + " " +
            strconv.Itoa(responseId)))
    }
}
