package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "strings"
)

type apiHandler struct {
    ApiModel apiModel
    request http.Request
    response http.ResponseWriter
}

// Load up the internal API model based on the provided API name
func (self * apiHandler) withApiName(apiName string) bool {
    database := database{}
    database.connect()
    // @TODO Would be better if this was consolidated
    apiRows := database.getApis()
    found := false
    for apiRows.Next() {
        model := apiModel{}
        model.loadFromRow(apiRows)
        if(model.Name == apiName) {
            model.loadMessages()
            self.ApiModel = model
            found = true
            break
        }
    }
    return found
}

func (self *apiHandler) respond () {
    body, error := ioutil.ReadAll(io.LimitReader(self.request.Body, 1048576))
    if (error != nil) {
        fmt.Fprintln(self.response, error)
        return
    }
    var messagePosition int
    for key, messageModel := range self.ApiModel.Messages {
        if(strings.Contains(string(body), messageModel.Identifier)) {
            messagePosition = key
            break
        }
    }
    if(messagePosition == 0) {
        fmt.Fprintln(self.response, "Message not found")
        return
    }
    responsesModel := self.ApiModel.Messages[messagePosition].getResponse(string(body))
    fmt.Fprintln(self.response, responsesModel.Template)
}
