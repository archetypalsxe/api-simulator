package main

import (
    "log"
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
    log.Fatal("No messages matched provided message")

    // Pointless, but Golang requires it
    model := messagesModel{}
    return model
}
