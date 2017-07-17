package main

import (
    "testing"
)

func TestDetermineMessageExactMatch (test *testing.T) {
    apiModel := apiModel{}
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 1, Identifier: "something different"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 2, Identifier: "testing Find this another"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 3, Identifier: "something else different"})

    messageParser := messageParser{ApiModel: apiModel}
    foundMessage := messageParser.determineMessage("find This")

    if(foundMessage.Id != 2) {
        test.Error("Correct message was not returned")
    }
}
