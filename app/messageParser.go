package main

import (
    "log"
    "strconv"
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
func (self * messageParser) parseMessage(body string, apiModel apiModel, messagePosition int) []messageFieldsModel {
    // @TODO Remove
    log.Println(apiModel.BeginningEscape)
    //"garabageField 123VALUEClosingmore garbage"
    //messagesModel{Id: 1, Identifier: "something different", Template: "**Field 123<<123>>Closing**"}
    // @TODO May not need the wild card

    for len(body) > 0 {
        // @TODO Remove
        log.Println(body)
        firstStart := strings.Index(body, apiModel.BeginningEscape)
        firstEnd := strings.Index(body, apiModel.EndingEscape)
        log.Println(strconv.Itoa(firstStart))
        log.Println(strconv.Itoa(firstEnd))
        if(firstStart > -1) {
            if(firstEnd < 0) {
                log.Fatal("We have a start without an end!")
            }
            // @TODO Get the data between
            body = body[(firstEnd+len(apiModel.EndingEscape)):len(body)]
        } else {
            body = ""
        }
    }

    // @TODO Remove
    log.Fatal("Testing")

    // @TODO Remove
    /*
    positions := self.getPositions(body, apiModel.BeginningEscape)
    for _, position := range positions {
        positionString := strconv.Itoa(position)
        log.Printf(positionString)
    }

    wildCards := self.getPositions(body, apiModel.WildCard)
    for _, wildCardPosition := range wildCards {
        wildCardPositionString := strconv.Itoa(wildCardPosition)
        log.Printf(wildCardPositionString)
    }
    log.Fatal("Testing")

    /**
     * Parse the message into parts, ie value, start, end, identifier
     */
    var messageFields []messageFieldsModel
    return messageFields
}

/*
func (self * messageParser) getPositions(haystack string, needle string) []int {
    var beginnings []int

log.Println(haystack)
log.Println(needle)
    for(strings.Index(haystack, needle) > -1) {
log.Println(haystack)
        beginnings = append(beginnings, strings.Index(haystack, needle))
        strings.Replace(haystack, needle, "", 1)
    }

    return beginnings
}
*/
