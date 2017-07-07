package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "regexp"
    "strconv"
    "strings"

    // Third party code for routing
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/settings", Settings)
    router.HandleFunc("/updateSettings", UpdateSettings)
    router.HandleFunc("/example/{id}", ExampleId)
    router.HandleFunc("/worldspan", Worldspan)
    router.HandleFunc("/api/{apiName}", ApiCall)

    // File servers
    path := os.Getenv("GOPATH") + "/src/api-simulator/"
    router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
        http.FileServer(http.Dir(path + "css"))))
    router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
        http.FileServer(http.Dir(path + "/js"))))

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Root directory of the API simulator\n")
}

func Settings(response http.ResponseWriter, request *http.Request) {
    settingsPage := settingsPage{response: response}
    settingsPage.respond()
}

// Handles Ajax requests for updating various settings
func UpdateSettings(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()

    switch request.FormValue("action") {
        case "saveApi":
            saveApiFromForm(request, response)
        case "saveMessage":
            saveMessageFromForm(request, response)
        case "saveResponse":
            saveResponseFromForm(request, response)
        case "saveNewField":
            saveNewFieldFromForm(request, response)
        case "updateField":
            updateFieldFromForm(request, response)
        default:
            log.Fatal("Invalid action requested")
    }
}

// Save a provided field that was entered on the form
func updateFieldFromForm(request *http.Request, response http.ResponseWriter) {
    id := request.FormValue("id")
    value := request.FormValue("value")

    // Get the ID from the id
    regExNums := regexp.MustCompile("[0-9]+")
    numbers := regExNums.FindAllString(id, -1)
    var databaseId string
    if(len(numbers) > 0) {
        databaseId = numbers[len(numbers)-1]
    }

    // Get the field name from the id
    regExLetters := regexp.MustCompile("[A-Za-z]+")
    letters := regExLetters.FindAllString(id, 1)
    if(len(letters) < 1) {
        log.Fatal("Did not have any letters in the id field")
    }
    fieldName := letters[0]

    switch fieldName {
        case "responseField":
            model := responsesModel{}
            id, _ := strconv.Atoi(databaseId)
            model.loadFromId(id)
            model.Template = value
            model.save()
        case "identifierField", "messageField":
            model := messagesModel{}
            id, _ := strconv.Atoi(databaseId)
            model.loadFromId(id)
            switch fieldName {
                case "identifierField":
                    model.Identifier = value
                case "messageField":
                    model.Template = value
            }
            database := database{}
            database.connect()
            database.updateMessages(model)
        case "apiNameField", "beginningEscapeField", "endingEscapeField", "apiWildCard":
            apiModel := apiModel{}
            apiModel.loadFromId(databaseId)
            switch fieldName {
                case "apiNameField":
                    apiModel.Name = value
                case "beginningEscapeField":
                    apiModel.BeginningEscape = value
                case "endingEscapeField":
                    apiModel.EndingEscape = value
                case "apiWildCard":
                    apiModel.WildCard = value
            }
            database := database{}
            database.connect()
            database.updateApi(apiModel)
        default:
            log.Fatal("Unknown field provided")
    }
}

// Save an API that was entered on the settings page
func saveApiFromForm(request *http.Request, response http.ResponseWriter) {
    model := apiModel{Name: request.FormValue("apiName"),
        BeginningEscape: request.FormValue("beginningEscape"),
        EndingEscape: request.FormValue("endingEscape"),
        WildCard: request.FormValue("wildCard")}
    database := database{}
    database.connect()
    result := database.insertApi(model)
    ajaxResponse := ajaxResponse{Status: result, Error: "None"}
    json.NewEncoder(response).Encode(ajaxResponse)
}

// Saves a messgae that was entered on the settings page
func saveMessageFromForm(request *http.Request, response http.ResponseWriter) {
    apiId, _ := strconv.Atoi(request.FormValue("apiId"))
    model := messagesModel{ApiId: apiId,
        Identifier: request.FormValue("identifier")}
    database := database{}
    database.connect()
    result := database.insertMessage(model)
    ajaxResponse := ajaxResponse{Status: result, Error: "None"}
    json.NewEncoder(response).Encode(ajaxResponse)
}

// Save the response that was received from the submitted form
func saveResponseFromForm(request *http.Request, response http.ResponseWriter) {
    messageId, _ := strconv.Atoi(request.FormValue("messageId"))
    isDefault, _ := strconv.ParseBool(request.FormValue("isDefault"))
    model := responsesModel {MessageId: messageId,
        Template: request.FormValue("response"),
        Default: isDefault,
        Condition: request.FormValue("condition")}
    database := database{}
    database.connect()
    result := database.insertResponse(model)
    ajaxResponse := ajaxResponse{Status: result, Error: "None"}
    json.NewEncoder(response).Encode(ajaxResponse)
}

// Save a new message field that was received from the settings page
func saveNewFieldFromForm(request *http.Request, response http.ResponseWriter) {
    messageId, _ := strconv.Atoi(request.FormValue("id"))
    // Currently we are assuming that it's a message
    //fieldType := request.FormValue("type")
    fieldValue := request.FormValue("value")
    model := messageFieldsModel {MessageId: messageId,
        FieldName: fieldValue}
    database := database{}
    database.connect()
    result, id := database.insertMessageField(model)
    ajaxResponse := ajaxResponse{Status: result, Error: "None", Id: id}
    json.NewEncoder(response).Encode(ajaxResponse)

}

func ExampleId(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    identifier := strings.TrimSpace(providedVars["id"])
    fmt.Fprintln(response, "Provided ID: ", identifier)
}

func ApiCall(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    apiName := strings.TrimSpace(providedVars["apiName"])
    apiHandler := apiHandler{request: *request, response: response}
    if(!apiHandler.withApiName(apiName)) {
        fmt.Fprintln(response, "Invalid API name provided: "+ apiName)
        return
    }
    apiHandler.respond()
}

func Worldspan(response http.ResponseWriter, request *http.Request) {
    worldspan := worldspanConnection{response: response, request: request}
    worldspan.respond()
}

