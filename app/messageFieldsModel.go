package main

import (
    "database/sql"
)

type messageFieldsModel struct {
    Id int
    MessageId int
    FieldName string
}

// Load up this object with a provided row from the database
func (self * messageFieldsModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.MessageId, &self.FieldName)
}
