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
    panic("No messages matched provided request!")

    // Pointless, but Golang requires it
    model := messagesModel{}
    return model
}

// Parse the provided message for any necessary dynamic fields
func (self * messageParser) parseMessage(body string, apiModel apiModel, messagePosition int) []messageFieldsModel {
    template := apiModel.Messages[messagePosition].Template

    // @TODO Remove
    log.Println(body)

    for len(template) > 0 {
        // @TODO Remove
        log.Println(template)
        firstWildCard := strings.Index(template, apiModel.WildCard)
        firstBeginningEscape := strings.Index(template, apiModel.BeginningEscape)
        firstEndingEscape := strings.Index(template, apiModel.EndingEscape)
        if(firstWildCard < firstBeginningEscape) {
            template = template[(firstWildCard+len(apiModel.WildCard)):len(template)]
        } else {
            if(firstBeginningEscape > firstEndingEscape) {
                log.Fatal("Ending escape is before the beginning escape")
            }
            preField := template[0:firstBeginningEscape]
            fieldName := template[(firstBeginningEscape+len(apiModel.BeginningEscape)):firstEndingEscape]
            modifiedTemplate := template[firstEndingEscape+len(apiModel.EndingEscape):len(template)]
            nextWildCard := strings.Index(modifiedTemplate, apiModel.WildCard)
            nextEscape := strings.Index(modifiedTemplate, apiModel.BeginningEscape)
            var nextField int = -1
            var fieldWidth int = 0
            if(nextWildCard > nextEscape && nextEscape >= 0) {
                nextField = nextEscape
                // @TODO Doing this 84732473210473201 times
                fieldWidth = len(apiModel.BeginningEscape)
            } else {
                nextField = nextWildCard
                fieldWidth = len(apiModel.WildCard)
            }
            if(nextField < 0) {
                nextField = len(template)
                fieldWidth = 0
            }
            postField := modifiedTemplate[0:nextField]
            template = modifiedTemplate[nextField+fieldWidth:len(modifiedTemplate)]
            log.Println(preField)
            log.Println(fieldName)
            log.Println(postField)
            // Getting all the fields from the template, need to actually parse the message now
        }


        /*
        // @TODO Maybe remove everything below
        // @TODO Remove
        log.Println(body)
        //firstWildCard := strings.Index(body, apiModel.WildCard)
        // The body after the first wild card
        modifiedBody := body[(firstWildCard+len(apiModel.WildCard)):len(body)]
        secondWildCard := strings.Index(modifiedBody, apiModel.WildCard)
        firstStart := strings.Index(body, apiModel.BeginningEscape)
        firstEnd := strings.Index(body, apiModel.EndingEscape)
        log.Println(strconv.Itoa(firstWildCard))
        log.Println(strconv.Itoa(secondWildCard))
        log.Println(strconv.Itoa(firstStart))
        log.Println(strconv.Itoa(firstEnd))
        if(firstStart > -1) {
            if(firstEnd < 0) {
                log.Fatal("We have a start without an end!")
            }
            // Check to make sure that we don't have 2 wild cards before the first starting point
            if(firstStart > secondWildCard) {
                // @TODO Get the data between
                body = body[(firstEnd+len(apiModel.EndingEscape)):len(body)]
            } else {
                body = modifiedBody
            }
        } else {
            body = ""
        }
        */
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
