package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
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
        default:
            panic("Invalid action requested")
    }
}

// Save an API that was entered on the settings page
func saveApiFromForm(request *http.Request, response http.ResponseWriter) {
    model := apiModel{Name: request.FormValue("apiName"),
        BeginningEscape: request.FormValue("beginningEscape"),
        EndingEscape: request.FormValue("endingEscape")}
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
        Identifier: request.FormValue("identifier"),
        ResponseTemplate: request.FormValue("response")}
    database := database{}
    database.connect()
    result := database.insertMessage(model)
    ajaxResponse := ajaxResponse{Status: result, Error: "None"}
    json.NewEncoder(response).Encode(ajaxResponse)
}

func ExampleId(response http.ResponseWriter, request *http.Request) {
    providedVars := mux.Vars(request)
    identifier := strings.TrimSpace(providedVars["id"])
    fmt.Fprintln(response, "Provided ID: ", identifier)
}

func Worldspan(response http.ResponseWriter, request *http.Request) {
    worldspan := worldspanConnection{response: response, request: request}
    worldspan.respond()
}

