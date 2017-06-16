package main

import (
    "database/sql"
)

type messagesModel struct {
    Id int
    ApiId int
    Identifier string
    Responses []responsesModel
}

func (self * messagesModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.ApiId, &self.Identifier)
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
