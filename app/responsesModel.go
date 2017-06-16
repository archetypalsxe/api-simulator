package main

import (
    "database/sql"
)

type responsesModel struct {
    Id int
    Template string
    Default bool
    condition string
}

func (self * responsesModel) loadFromRow(row *sql.Rows) {
    row.Scan(&self.Id, &self.Template, &self.Default, &self.condition)
}
