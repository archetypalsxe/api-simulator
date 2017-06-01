package main

type apiModel struct {
    Id int
    Name string
    BeginningEscape string
    EndingEscape string
    Messages []messagesModel
}

/*
    Load up all the messages for this API model, assumes that the Id for
    this is set, or else we don't know where in the database to search
*/
func (self * apiModel) loadMessages() {
    var id int
    var apiId int
    var identifier string
    var responseId int

    database := database{}
    database.connect()
    rows := database.getMessagesForApi(self.Id)
    for rows.Next() {
        rows.Scan(&id, &apiId, &identifier, &responseId)
        model := messagesModel{Id: id, ApiId: apiId, Identifier:identifier,
            ResponseId: responseId}
        self.Messages = append(self.Messages, model)
    }
}
