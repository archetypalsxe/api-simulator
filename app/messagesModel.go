package main

import (
    "database/sql"
    "log"
)

type messagesModel struct {
    Id int
    ApiId int
    Identifier string
    Template string
    Fields []messageFieldsModel
    Responses []responsesModel
}

func (self * messagesModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.ApiId, &self.Identifier, &self.Template)
}

// Load up all the message fields that are related to this message
func (self * messagesModel) loadFields() {
    database := database{}
    database.connect()
    rows := database.getFieldsForMessage(self.Id)
    for rows.Next() {
        model := messageFieldsModel{}
        model.loadFromRow(rows)
        self.Fields = append(self.Fields, model)
    }
}

// Load up all the responses that are associated to this message
func (self * messagesModel) loadResponses() {
    database := database{}
    database.connect()
    rows := database.getResponsesForMessage(self.Id)
    for rows.Next() {
        model := responsesModel{}
        model.loadFromRow(rows)
        self.Responses = append(self.Responses, model)
    }
}

func (self * messagesModel) loadFromId(id int) {
    database := database{}
    database.connect()
    rows := database.getMessagesById(id)
    for rows.Next() {
        self.loadFromRow(rows)
    }
}

// Determine which response should be used and return the appropriate response
func (self * messagesModel) getResponse(body string) responsesModel {
    if(len(self.Responses) < 1) {
        log.Fatal("Not able to find response")
    }
    // @TODO Determine which response to use if there are multiple
    return self.Responses[0]
}
