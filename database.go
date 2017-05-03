package main

import (
    "database/sql"
    "log"

    //SQLite
    _ "github.com/mattn/go-sqlite3"
);

type database struct {
}

func (self *database) test() {
    database, error := sql.Open("sqlite3", "./settings.db");
    self.handleError(error)
    statement, error := database.Prepare(
        "CREATE TABLE IF NOT EXISTS testing (id INTEGER PRIMARY KEY)")
    self.handleError(error)
    statement.Exec()
}

func (self *database) handleError(error error) {
    if (error != nil) {
        log.Fatal(error);
    }
}
