package main

import (
    "database/sql"
    "log"

    //SQLite
    _ "github.com/mattn/go-sqlite3"
);

type database struct {
    connection *sql.DB
}

func (self *database) connect() {
    connection, error := sql.Open("sqlite3", "./settings.db")
    self.handleError(error)
    self.connection = connection
    self.initializeDatabase()
}

func (self *database) getApis() *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Apis;")
    self.handleError(error)
    return rows
}

func (self *database) getMessages() *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Messages;")
    self.handleError(error)
    return rows
}

func (self *database) insertData() {
    self.runQuery("DELETE FROM Apis;")
    self.runQuery("INSERT INTO Apis (name) VALUES ('worldspan');")
    self.runQuery("DELETE FROM Messages;")
    self.runQuery("INSERT INTO Messages (apiId, identifier, responseId) "+
        "VALUES ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<PSC5>', "+
            "1"+
        ")")
}

func (self *database) initializeDatabase() {
    self.runQuery("CREATE TABLE IF NOT EXISTS Apis "+
        "(id INTEGER PRIMARY KEY, name text)")
    self.runQuery("CREATE TABLE IF NOT EXISTS Messages ("+
            "id INTEGER PRIMARY KEY,"+
            "apiId INTEGER,"+
            "identifier TEXT,"+
            "responseId INTEGER"+
        ")")
    self.runQuery("CREATE TABLE IF NOT EXISTS Responses ("+
            "id INTEGER PRIMARY KEY,"+
            "template TEXT"+
        ")")
}

func (self *database) runQuery(query string) {
    statement, error := self.connection.Prepare(query)
    self.handleError(error)
    _, statementError := statement.Exec()
    self.handleError(statementError)
}

func (self *database) handleError(error error) {
    if (error != nil) {
        log.Fatal(error)
    }
}
