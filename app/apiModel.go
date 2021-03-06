package main

import (
    "database/sql"
)

type apiModel struct {
    Id int
    Name string
    BeginningEscape string
    EndingEscape string
    WildCard string
    Messages []messagesModel
}

/*
    Load up all the messages for this API model, assumes that the Id for
    this is set, or else we don't know where in the database to search
*/
func (self * apiModel) loadMessages() {
    database := database{}
    database.connect()
    rows := database.getMessagesForApi(self.Id)
    for rows.Next() {
        model := messagesModel{}
        model.loadFromRow(rows)
        model.loadFields()
        model.loadResponses()
        self.Messages = append(self.Messages, model)
    }
}

// Load up this object from a provided row
func (self * apiModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.Name, &self.BeginningEscape,
        &self.EndingEscape, &self.WildCard)
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
