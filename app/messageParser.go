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

    var messageFields []messageFieldsModel

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
            // @TODO Remove debugging code
            log.Println(preField)
            log.Println(fieldName)
            log.Println(postField)

            messageFieldsModel := messageFieldsModel{MessageId: apiModel.Messages[messagePosition].Id,
                FieldName: fieldName}

            // Getting all the fields from the template, need to actually parse the message now
            body = body[strings.Index(body, preField) + len(preField):len(body)]
            fieldValue := body[0:strings.Index(body, postField)]
            messageFieldsModel.Value = fieldValue
            messageFieldsModel.display()
            messageFields = append(messageFields, messageFieldsModel)
        }
    }

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
