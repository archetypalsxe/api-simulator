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

func (self * messagesModel) loadFromId(id int) {
    database := database{}
    database.connect()
    rows := database.getMessagesById(id)
    for rows.Next() {
        self.loadFromRow(rows)
    }
}
