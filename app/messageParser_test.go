package main

import (
    "testing"
)

func setup() apiModel {
    apiModel := apiModel{BeginningEscape: "<<", EndingEscape: ">>", WildCard: "**"}

    messagesModel1 := messagesModel{Id: 1, Identifier: "something different", Template: "**Field 123<<123>>Closing**"}
    messagesModel2 := messagesModel{Id: 2, Identifier: "testing Find this another", Template: "**12<<12>>/12**13<<13>>/13**"}
    messagesModel3 := messagesModel{Id: 3, Identifier: "something else different", Template: "<<14>>end**start<<15>>"}

    apiModel.Messages = append(apiModel.Messages, messagesModel1)
    apiModel.Messages = append(apiModel.Messages, messagesModel2)
    apiModel.Messages = append(apiModel.Messages, messagesModel3)

    return apiModel
}

// Test that we're able to pull a message if the case doesn't match
func TestDetermineMessageMatch (test *testing.T) {

    apiModel := setup()

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

    apiModel := setup()

    messageParser := messageParser{ApiModel: apiModel}
    messageParser.determineMessage("Shouldn't be found")

    test.Error("This spot shouldn't be reached")
}

// Test the parse message function to make sure we're getting matches
func TestParseMessage (test *testing.T) {
    apiModel := setup()
    messageParser := messageParser{ApiModel: apiModel}
    fields := messageParser.parseMessage("garbageField 123<<VALUE>>Closingmore garbage",
        apiModel, 0)
    if len(fields) < 1 {
        test.Error("Empty array returned")
    }
    for _, field := range fields {
        if field.Value != "VALUE" {
            test.Error("Did not find correct field")
        }
    }
}

// Test the parse message function to make sure we're getting matches. 2 fields
func TestParseMessageTwo (test *testing.T) {
    apiModel := setup()
    messageParser := messageParser{ApiModel: apiModel}
    fields := messageParser.parseMessage("garabage12THE VALUE/12some mo'\"fdsfdsatext13second value/13MORE TEXT",
        apiModel, 2)
    if len(fields) < 1 {
        test.Error("Empty array returned")
    }
    for position, field := range fields {
        if position == 0 && field.Value != "THE VALUE" {
            test.Error("Did not find correct field")
        }
        if position == 1 && field.Value != "second value" {
            test.Error("Did not find correct field")
        }
    }
}

// Test the parse message function to make sure we're getting matches. Fields at beginning and end
func TestParseMessageBeginningEnd (test *testing.T) {
    apiModel := setup()
    messageParser := messageParser{ApiModel: apiModel}
    fields := messageParser.parseMessage("the valueend random junk startlast value",
        apiModel, 2)
    if len(fields) < 1 {
        test.Error("Empty array returned")
    }
    for position, field := range fields {
        if position == 1 && field.Value != "the value" {
            test.Error("Did not find correct field")
        }
        if position == 1 && field.Value != "last value" {
            test.Error("Did not find correct field")
        }
    }
}
