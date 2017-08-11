package main

import (
    "database/sql"
    "log"
)

type messageFieldsModel struct {
    Id int
    MessageId int
    FieldName string
    Value string
}

// Load up this object with a provided row from the database
func (self * messageFieldsModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.MessageId, &self.FieldName)
}

// Dump this object out to the console for testing
func (self * messageFieldsModel) display() {
    log.Println("Displaying Message Fields Model:");
    log.Println(self.Id)
    log.Println(self.MessageId)
    log.Println(self.FieldName)
    log.Println(self.Value)
    log.Println("End of Message Fields Model");
}
