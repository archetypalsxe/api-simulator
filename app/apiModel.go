package main

import (
    "database/sql"
)

type apiModel struct {
    Id int
    Name string
    BeginningEscape string
    EndingEscape string
    Messages []messagesModel
}

/*
    Load up all the messages for this API model, assumes that the Id for
    this is set, or else we don't know where in the database to search
*/
func (self * apiModel) loadMessages() {
    var id int
    var apiId int
    var identifier string
    var responseId int
    var responseTemplate string

    database := database{}
    database.connect()
    rows := database.getMessagesForApi(self.Id)
    for rows.Next() {
        rows.Scan(&id, &apiId, &identifier, &responseId)
        responseMessages := database.getResponseMessage(responseId)
        for responseMessages.Next() {
            responseMessages.Scan(&responseId, &responseTemplate)
        }
        model := messagesModel{Id: id, ApiId: apiId, Identifier:identifier,
            ResponseId: responseId, ResponseTemplate: responseTemplate}
        self.Messages = append(self.Messages, model)
    }
}

func (self * apiModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.Name, &self.BeginningEscape, &self.EndingEscape)
}

// Load up an API model (self) with the provided ID
func (self * apiModel) loadFromId(id string) {
    database := database{}
    database.connect()
    rows := database.getApi(id)
    for rows.Next() {
        self.loadFromRow(rows)
    }
}
