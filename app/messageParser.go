package main

import (
    "strings"
)

type messageParser struct {
    ApiModel apiModel
}

// Determines which message was provided
func (self * messageParser) determineMessage(body string) messagesModel {
    for _, messageModel := range self.ApiModel.Messages {
        if(strings.Contains(strings.ToLower(messageModel.Identifier), strings.ToLower(body))) {
            return messageModel
        }
    }
    panic("No messages matched provided request!")

    // Pointless, but Golang requires it
    model := messagesModel{}
    return model
}

// Parse the provided message for any necessary dynamic fields
func (self * messageParser) parseMessage(body string, message messagesModel) []messageFieldsModel {
    var messageFields []messageFieldsModel
    return messageFields
}
