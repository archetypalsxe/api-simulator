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
    database := database{}
    database.connect()
    rows := database.getMessagesForApi(self.Id)
    for rows.Next() {
    }
}
