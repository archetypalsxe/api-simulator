package main

import (
    "testing"
)

// Test that we're able to pull a message if the case doesn't match
func TestDetermineMessageMatch (test *testing.T) {
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

// Test that we're able to pull a message that has symbols in it
func TestDetermineMessageMatchSymbols (test *testing.T) {
    apiModel := apiModel{}
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 1, Identifier: "som'ethi%ng dif&fer'hent"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 2, Identifier: "tes<\"t'ing Find this anoth'>er"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 3, Identifier: "som\"ething e\"&%><\"lse different"})

    messageParser := messageParser{ApiModel: apiModel}
    foundMessage := messageParser.determineMessage("\"&%><\"")

    if(foundMessage.Id != 3) {
        test.Error("Correct message was not returned")
    }
}

// Test that we panic when the message can't be found
func TestPanicMessage (test *testing.T) {
    defer func() {
        if recovery := recover(); recovery != nil {
            return
        }
    }()

    apiModel := apiModel{}
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 1, Identifier: "something different"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 2, Identifier: "testing Find this another"})
    apiModel.Messages = append(apiModel.Messages, messagesModel{Id: 3, Identifier: "something else different"})

    messageParser := messageParser{ApiModel: apiModel}
    messageParser.determineMessage("Shouldn't be found")

    test.Error("This spot shouldn't be reached")
}
