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
    log.Println("****PARSING MESSAGE****")

    var messageFields []messageFieldsModel

    template := apiModel.Messages[messagePosition].Template

    // @TODO Remove
    log.Print("Body:" );
    log.Print(body)

    for len(template) > 0 {
        // @TODO Remove
        log.Print("Template: ");
        log.Print(template)
        firstWildCard := strings.Index(template, apiModel.WildCard)
        firstBeginningEscape := strings.Index(template, apiModel.BeginningEscape)
        firstEndingEscape := strings.Index(template, apiModel.EndingEscape)
log.Print(firstWildCard)
log.Print(firstBeginningEscape)
        if(firstWildCard < firstBeginningEscape && firstWildCard != -1) {
log.Print("Line 46")
            template = template[(firstWildCard+len(apiModel.WildCard)):len(template)]
        } else {
            if(firstBeginningEscape > firstEndingEscape) {
                log.Fatal("Ending escape is before the beginning escape")
            }
            preField := template[0:firstBeginningEscape]
            fieldName := template[(firstBeginningEscape+len(apiModel.BeginningEscape)):firstEndingEscape]
            modifiedTemplate := template[firstEndingEscape+len(apiModel.EndingEscape):len(template)]
log.Print("Modified template:")
log.Print(modifiedTemplate)
            nextWildCard := strings.Index(modifiedTemplate, apiModel.WildCard)
            nextEscape := strings.Index(modifiedTemplate, apiModel.BeginningEscape)
            var nextField int = -1
            var fieldWidth int = 0
log.Print("Next wild card: ")
log.Print(nextWildCard)
log.Print("Next escape:")
log.Print(nextEscape)
            if(nextWildCard > nextEscape && nextEscape >= 0) {
log.Print("Line 59")
                nextField = nextEscape
                // @TODO Doing this 84732473210473201 times
                fieldWidth = len(apiModel.BeginningEscape)
            } else {
log.Print("Line 64")
                nextField = nextWildCard
                fieldWidth = len(apiModel.WildCard)
            }
            if(nextField < 0) {
log.Print("Line 69")
                if(nextEscape < 0) {
log.Print("Line 80")
                    nextField = 0
                    fieldWidth = 0
                } else {
                    nextField = len(template)
                    fieldWidth = 0
                }
            }
log.Print(nextField)
log.Print(fieldWidth)
            postField := modifiedTemplate[0:nextField]
            template = modifiedTemplate[nextField+fieldWidth:len(modifiedTemplate)]
            // @TODO Remove debugging code
            log.Print("Pre Field: ")
            log.Print(preField)
            log.Print("Field Name: ")
            log.Print(fieldName)
            log.Print("Post Field: ")
            log.Print(postField)

            messageFieldsModel := messageFieldsModel{MessageId: apiModel.Messages[messagePosition].Id,
                FieldName: fieldName}

            // Getting all the fields from the template, need to actually parse the message now
            body = body[strings.Index(body, preField) + len(preField):len(body)]

            var fieldValue string

            if(postField == "") {
                fieldValue = body[0:len(body)]
            } else {
                fieldValue = body[0:strings.Index(body, postField)]
            }
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
