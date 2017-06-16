package main

import (
    "database/sql"
    "io/ioutil"
    "log"
    "os"
    "strconv"

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
    //self.initializeDatabase()
}

func (self *database) getApis() *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Apis;")
    self.handleError(error)
    return rows
}

func (self *database) getApi(apiId string) *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Apis WHERE "+
        "id = '"+ apiId + "';")
    self.handleError(error)
    return rows
}

func (self *database) getMessages() *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Messages;")
    self.handleError(error)
    return rows
}

func (self *database) getMessagesById(messageId int) *sql.Rows {
    query := "SELECT * FROM Messages WHERE id = "+
        strconv.Itoa(messageId) +";"
    rows, error := self.connection.Query(query)
    self.handleError(error)
    return rows
}

func (self *database) getMessagesForApi(apiId int) *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Messages WHERE "+
        "apiId = '"+ strconv.Itoa(apiId) +"';")
    self.handleError(error)
    return rows
}

func (self *database) getResponseMessage(responseId int) *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Responses WHERE "+
        "id = '"+ strconv.Itoa(responseId) +"';")
    self.handleError(error)
    return rows
}

func (self *database) getResponses() *sql.Rows {
    rows, error := self.connection.Query("SELECT * FROM Responses;")
    self.handleError(error)
    return rows
}

func (self *database) updateApi(apiModel apiModel) bool {
    query := "UPDATE Apis SET name = '"+ apiModel.Name +
        "', beginningEscape = '"+ apiModel.BeginningEscape +
        "', endingEscape = '"+ apiModel.EndingEscape +
        "' WHERE id = "+ strconv.Itoa(apiModel.Id) +";"
    result := self.runQuery(query)
    rowsAffected, _ := result.RowsAffected()
    return rowsAffected > 0
}

// Updating a provided message in the database
func (self *database) updateMessages(model messagesModel) bool {
    query := "UPDATE Messages SET identifier = '"+ model.Identifier +
        "';"
    result := self.runQuery(query)
    rowsAffected, _ := result.RowsAffected()
    return rowsAffected > 0
}

/// Insert the provided API into the database
func (self *database) insertApi(apiModel apiModel) bool {
    query := "INSERT INTO Apis (name, beginningEscape, endingEscape) "+
        "VALUES ('"+ apiModel.Name +"', '"+ apiModel.BeginningEscape +
        "', '"+ apiModel.EndingEscape +"');"
    result := self.runQuery(query)
    rowsAffected, _ := result.RowsAffected()
    return rowsAffected > 0
}

// Insert the provided message (and response(s)) in the database
// @TODO Needs to be redone
func (self *database) insertMessage(messagesModel messagesModel) bool {
    /*
    query := "INSERT INTO Responses (template) VALUES "+
        "('"+ messagesModel.ResponseTemplate +"')"
    response := self.runQuery(query)
    responseId, _ := response.LastInsertId()
    // @TODO This could become an issue in the future, converting int64 to int
    responseIdString := strconv.Itoa(int(responseId))
    apiIdString := strconv.Itoa(messagesModel.ApiId)
    insertResult := self.runQuery("INSERT INTO Messages (apiId, identifier, responseId) "+
        "VALUES ("+
            "'"+ apiIdString +"', "+
            "'"+ messagesModel.Identifier +"', "+
            responseIdString +
        ")")
    rowsAffected, _ := insertResult.RowsAffected()
    return rowsAffected > 0
    */
    return true
}

func (self *database) insertData() {
    // Insert the APIs
    self.runQuery("DELETE FROM Apis;")
    self.runQuery("INSERT INTO Apis (name, beginningEscape, endingEscape) "+
        "VALUES ('worldspan', '!--', '--!');")

    // Insert the responses
    // @TODO Should be using os.PathSeparator
    dataPath := os.Getenv("GOPATH") + "/src/api-simulator/data/";
    self.runQuery("DELETE FROM Responses")
    self.runQuery("INSERT INTO Responses (template) VALUES"+
        "('"+self.getFileContents(dataPath+"testCardAuthorization") +"'),"+
        "('"+self.getFileContents(dataPath+"testDisplayPnrResponse") +"'),"+
        "('"+self.getFileContents(dataPath+"testFinished") +"'),"+
        "('"+self.getFileContents(dataPath+"testNativeDisplayPnrResponse") +"'),"+
        "('"+self.getFileContents(dataPath+"testPowerShopperResponse") +"'),"+
        "('"+self.getFileContents(dataPath+"testPricingResponse") +"'),"+
        "('"+self.getFileContents(dataPath+"testTicketingResponse") +"'),"+
        "('"+self.getFileContents(dataPath+"testUpdatePnrResponse") +"')")

    // Insert the messages
    self.runQuery("DELETE FROM Messages;")
    self.runQuery("INSERT INTO Messages (apiId, identifier) "+
        "VALUES ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<HOS_CMD>CK/'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<DPC8>'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<HOS_RSP_SCR>F</HOS_RSP_SCR>'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<HOS_CMD>*'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<PSC5>'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<BPC9>'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<HOS_CMD>EZEI#$*'"+
        "), ("+
            "(SELECT id FROM Apis WHERE name = 'worldspan'), "+
            "'<UPC7>'"+
        ")")
    self.runQuery("INSERT INTO MessagesResponsesMap (messagesId, responsesId) "+
        "VALUES "+
            "(1, 1),(2,2),(3,3),(4,4),(5,5),(6,6),(7,7),(8,8);")
}

func (self *database) initializeDatabase() {
    self.runQuery("DROP TABLE IF EXISTS Apis;");
    self.runQuery("DROP TABLE IF EXISTS Messages;");
    self.runQuery("DROP TABLE IF EXISTS Responses;");
    self.runQuery("DROP TABLE IF EXISTS MessagesResponsesMap;");
    self.runQuery("CREATE TABLE IF NOT EXISTS Apis ("+
        "id INTEGER PRIMARY KEY,"+
        "name text,"+
        "beginningEscape text,"+
        "endingEscape text"+
   ")")
    self.runQuery("CREATE TABLE IF NOT EXISTS Messages ("+
            "id INTEGER PRIMARY KEY,"+
            "apiId INTEGER,"+
            "identifier TEXT"+
        ")")
    self.runQuery("CREATE TABLE IF NOT EXISTS Responses ("+
            "id INTEGER PRIMARY KEY,"+
            "template TEXT"+
        ")")
    self.runQuery("CREATE TABLE IF NOT EXISTS  MessagesResponsesMap ("+
        "messagesId INTEGER NOT NULL,"+
        "responsesId INTEGER NOT NULL,"+
        "`default` INTEGER DEFAULT 1 NOT NULL,"+
        "condition TEXT,"+
        "PRIMARY KEY(messagesId, responsesId)"+
        ")")
}

func (self *database) runQuery(query string) sql.Result {
    statement, error := self.connection.Prepare(query)
    self.handleError(error)
    response, statementError := statement.Exec()
    self.handleError(statementError)
    return response
}

func (self *database) handleError(error error) {
    if (error != nil) {
        log.Fatal(error)
    }
}

func (self *database) getFileContents(responseFile string) string {
    data, error := ioutil.ReadFile(responseFile);
    if (error != nil) {
        log.Fatal(error);
    }
    return string(data)
}
