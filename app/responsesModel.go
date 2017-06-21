package main

import (
    "database/sql"
    "log"
)

type responsesModel struct {
    Id int
    Template string
    Default bool
    Condition string
    MessageId int
}

// Load up this object from the provided database row
func (self * responsesModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.Template, &self.Default, &self.Condition)
}

// Load up this object with the provided database ID
func (self * responsesModel) loadFromId(id int) {
    database := database{}
    database.connect()
    rows := database.getResponseMessage(id)
    for rows.Next() {
        self.loadFromRow(rows)
    }
    if self.Id == 0 {
        log.Fatal("Not able to find response in database")
    }
}

// Save this model into the database
func (self * responsesModel) save() {
    database := database{}
    database.connect()
    database.updateResponse(*self)
}
